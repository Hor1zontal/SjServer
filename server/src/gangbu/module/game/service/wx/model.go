package wx

type LoginResponse struct {
	OpenID 		string `json:"openid"`
	SessionKey 	string `json:"session_key"`
	UnionID 	string `json:"unionid"`
}

type ErrorResponse struct {
	ErrCode 	int `json:"errcode"`
	ErrMsg 		string `json:"errmsg"`
}