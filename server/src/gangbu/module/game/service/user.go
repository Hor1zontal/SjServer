package service

import (
	"aliens/log"
	"gangbu/constant"
	"gangbu/exception"
	"gangbu/module/game/config"
	"gangbu/module/game/db"

	"gangbu/module/game/service/myjwt"
	"gangbu/module/game/service/wx"
	"gangbu/module/game/words"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Login(code string, channel int32) *db.DBUser {
	var user *db.DBUser
	var isNew bool
	if channel == constant.CHANNEL_VISITOR || channel == constant.CHANNEL_WEB {
		nickname := words.RandomName()
		user, isNew = GetUserUsername("sj_" + code, "", nickname, channel)
	} else if channel == constant.CHANNEL_TOUTIAO || channel == constant.CHANNEL_WECHAT {
		user, isNew = GetUserByCode(code, channel)
	} else {
		exception.GameException(exception.PlatformUnknown)
	}
	user.UpdateActiveTime()
	user.LoginLog(isNew, channel)
	db.UpdateOne(user)
	return user
}

func GetUserUsername(username, openid, nickname string, channel int32) (*db.DBUser, bool) {
	user := db.GetUserInfo(username)
	var isNew bool
	if user == nil {
		//用户不存在
		userNew := db.NewUser(username,"", openid, "", nickname, channel)
		user = userNew
		isNew = true
	}
	return user, isNew
}

func GetUserByCode(code string, channel int32) (*db.DBUser, bool) {
	openid, _, err := wx.WxLogin(code, channel)
	if err != nil {
		exception.GameExceptionCustom("GetUserByCode-WxLogin", exception.WxLogin, err)
	}
	user, isNew := GetUserUsername("sj_" + openid, openid, "", channel)
	return user, isNew
}

func UpdateUser(uid int32, nickname, avatar string) {

	user := db.GetUserByUid(uid)
	if user == nil {
		exception.GameException(exception.UserNotFound)
	}
	if user.Nickname != nickname {
		if user.Channel == constant.CHANNEL_WEB || user.Channel == constant.CHANNEL_TOUTIAO {
			words.ValidateNickname(nickname)
		}
		user.Nickname = nickname
	}
	user.Avatar = avatar
	db.UpdateOne(user)
}

// 生成令牌
func GenerateToken(uid int32) string {
	j := myjwt.NewJWT(config.Server.JWTSecret)
	claims := myjwt.CustomClaims{
		UID:uid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),	// 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + config.Server.JWTExpired),	// 过期时间 一小时
			Issuer:    "aliens",					//签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		log.Error(err.Error())
		return ""
	}

	//log.Debug(token)
	return token
}

func DeleteUser(uid int32) {
	db.DeleteUser(uid)
	db.DeleteRole(uid)
}