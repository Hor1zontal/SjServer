package api

import (
	"aliens/common/character"
	"aliens/common/cipher"
	"aliens/log"
	"bytes"
	"encoding/base64"
	"gangbu/exception"
	"gangbu/module/game/config"
	"gangbu/module/game/service/myjwt"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/sortkeys"
	"io/ioutil"
	"time"
)

func JWTAuth(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			exception.GameException(exception.TokenIsNil)
		}
		//log.Debug("get token:%v", token)

		j := myjwt.NewJWT(config.Server.JWTSecret)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == myjwt.TokenExpired {
				exception.GameException(exception.TokenExpired)
			} else {
				log.Error(err.Error())
				exception.GameException(exception.TokenCheckError)
			}
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
}

func JWTRefresh(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		exception.GameException(exception.TokenIsNil)
	}
	log.Debug("get token:%v", token)
	j := myjwt.NewJWT(config.Server.JWTSecret)
	token, err := j.RefreshToken(token)
	if err != nil {
		exception.GameException(exception.TokenCheckError)
	}
	c.Header("token", token)
}

func SignCheck(c *gin.Context) {
	if !config.Server.IsSign {
		return
	}
	data, _ := ioutil.ReadAll(c.Request.Body)
	reqBody := ioutil.NopCloser(bytes.NewBuffer(data))
	c.Request.Body = reqBody
	c.Request.ParseForm()
	c.Request.Body = reqBody

	sign := c.Request.Header.Get("sign")
	if !isSignSuccess(c.Request.Form, sign) {
		exception.GameException(exception.SignError)
	}
}
func isSignSuccess(reqData map[string][]string, signData string) bool {

	str_timestamp := ""
	for _, value := range reqData["timestamp"] {
		str_timestamp += value
	}
	timestamp := character.StringToInt64(str_timestamp) / 1000
	tm_now := time.Now()
	tm := time.Unix(timestamp,0)

	duration:=  tm_now.Sub(tm)
	log.Debug("%v", duration.Seconds())
	//ss, _ :=time.ParseDuration("1s")
	if tm_now.Before(tm.Add(time.Duration(config.Server.ExpiredTime)*time.Second)) {
		//log.Debug("%v", time.Now().Sub(tm))
		log.Debug("before")
	} else {
		exception.GameException(exception.SignError)
	}

	var signText string
	var strKeys = []string{}
	for key := range reqData {
		strKeys = append(strKeys, key)
	}
	sortkeys.Strings(strKeys)

	for _, value := range strKeys {
		if value != "sign" {
			dataText := ""
			for _,value := range reqData[value] {
				dataText += value
			}
			if value == "nickname" {
				// 为了避免出现名字为emoji时与客户端签名不一致的情况进行base64加密
				dataText = base64.StdEncoding.EncodeToString([]byte(dataText))
			}
			signText += dataText
		}
	}
	//todo 加上签名的key
	// signText += conf.Server.SignKey
	log.Debug("signText:%v",signText)
	signResult := cipher.MD5Hash(signText)
	if signResult != signData {
		log.Debug("signResult %v : sign : %v", signResult, signData)
		return false
	}

	return true
}