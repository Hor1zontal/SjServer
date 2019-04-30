package api

import (
	"gangbu/exception"
	"gangbu/module/game/http/helper"
	"gangbu/module/game/service"
	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /users/ 获取用户信息
 * @apiGroup Users
 * @apiParam {Number} code 微信的code
 * @apiSuccess {string} token 用户的token
 * @apiSuccess {string} nickname 用户昵称
 * @apiSuccess {string} avatar 用户头像地址
 * @apiSuccess {Number} errcode 错误码
 * @apiSuccess {String} errmsg 错误信息
 */
type GetUsersReq struct {
	Code 		string  `form:"code" binding:"required"`
	Channel 	int32  `form:"channel" binding:"required"`
}

type GetUsersResp struct {
	UID   	 int32  `json:"uid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func Login(c *gin.Context) {

	req := &GetUsersReq{}
	//if err := c.ShouldBind(req); err != nil {
	//	log.Error(err.Error())
	//	exception.GameException(exception.InvalidParam)
	//}
	helper.CheckReq(c, req)

	//log.Info("bind success")
	//user := service.GetUserByCode(req.Code)
	user := service.Login(req.Code, req.Channel)
	token := service.GenerateToken(user.ID)

	resp := &GetUsersResp{}
	resp.UID = user.ID
	resp.Token = token
	resp.Nickname = user.Nickname
	resp.Avatar = user.Avatar

	helper.ResponseWithData(c, resp)
}


/**
 * @api {PUT} /users/ 更新用户头像昵称
 * @apiGroup Users
 * @apiParam {Number} uid 用户的uid
 * @apiParam {String} token 用户的token
 * @apiParam {String} nickname 用户的昵称
 * @apiParam {String} avatar 用户头像地址
 * @apiSuccess {Number} errcode 错误码
 * @apiSuccess {String} errmsg 错误信息
 */
type PutUsersReq struct {
	Nickname string `form:"nickname" binding:"required"`
	Avatar   string `form:"avatar" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	//tUser := helper.GetClaimUser(c)
	req := &PutUsersReq{}
	//if err := c.ShouldBind(req); err != nil {
	//	log.Error(err.Error())
	//	exception.GameException(exception.InvalidParam)
	//}
	helper.CheckReq(c,req)
	service.UpdateUser(helper.GetClaimUser(c).UID, req.Nickname, req.Avatar)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}

type DeleteUserReq struct {
	Uid int32 `form:"uid"`
}

func DeleteUser(c *gin.Context) {
	req := &DeleteUserReq{}
	helper.CheckReq(c, req)
	service.DeleteUser(req.Uid)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}