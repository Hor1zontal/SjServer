package exception

import "aliens/log"

type ErrorCode int32

/**
* @api {ERROR_CODE} 错误码
* @apiGroup ERROR
* @apiSuccess {Number} errcode 错误码
* @apiSuccess {String} errmsg 错误信息
* @apiSuccess (errcode) {Number} 10001 参数错误
* @apiSuccess (errcode) {Number} 10002 数据库操作异常
* @apiSuccess (errcode) {Number} 10003 内部异常
* @apiSuccess (errcode) {Number} 20001 微信登录异常
* @apiSuccess (errcode) {Number} 20002 无效的token
* @apiSuccess (errcode) {Number} 20003 未携带token
*/
const (
	CodeSuccess   ErrorCode = iota
	InvalidParam            = 10001 //参数错误
	DatabaseError           = 10002 //数据库操作异常
	InternalError           = 10003 //内部异常
	TimeParseError			= 10004 //时间格式错误
	SignError				= 10005 //签名错误

	WxLogin      			= 20001 //微信登录异常
	TokenInvalid 			= 20002 //无效的token
	TokenIsNil				= 20003 //未携带token
	TokenExpired			= 20004 //token过期
	TokenCheckError			= 20005 //token验证错误
	EnergyNotEnough			= 20006 //体力不足
	UpdateScoreError		= 20007 //上传分数错误
	GameDataNotFound		= 20008 //游戏数据未找到
	PropNotFound			= 20009 //游戏道具未找到
	PropNotEnough			= 20010 //道具不足
	RoleNotFound			= 20011 //用户角色未找到
	PlatformUnknown			= 20013 //平台未知
	NicknameInvalidWord		= 20014 //不合法的名字
	NicknameSensitiveWord	= 20015 //含有敏感词
	UserNotFound			= 20016 //用户未找到
	//ErrorCodeUserNotFound           = 20002 //用户未找到
)

func GameException(this ErrorCode) {
	panic(this)
}

var ErrorMapping = map[ErrorCode]string{
	CodeSuccess:	"success",
	InvalidParam:	"invalid param",
	DatabaseError:	"database error",
	InternalError:	"internal error",
	TimeParseError: "time parse error",
	SignError: "sign error",

	WxLogin:		"weChat login error",
	TokenInvalid:	"token invalid",
	TokenIsNil:		"token is nil",
	TokenExpired:	"token expired",
	TokenCheckError:"token check error",
	EnergyNotEnough:"energy not enough",
	UpdateScoreError:"you cheated",
	GameDataNotFound:"game data not found",
	PropNotFound:"prop not found",
	PropNotEnough: "prop not enough",
	RoleNotFound: "role not found",
	PlatformUnknown: "platform unknown",
	NicknameInvalidWord: "nickname invalid word",
	NicknameSensitiveWord: "nickname sensitive word",
	UserNotFound: "user not found",
}


func GameExceptionCustom(api string, code ErrorCode, err error) {
	log.Error(api + " -- code:%v --- error:%v", code, err)
	GameException(code)
}
