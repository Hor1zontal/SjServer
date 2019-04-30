package wx

import (
	"bytes"
	"gangbu/module/game/config"
	"encoding/json"
	"fmt"
	"gangbu/constant"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"net/http"
)

const (
	PARAM_APPID = "appid"
	PARAM_SECRET = "secret"
	PARAM_CODE = "js_code"
	PARAM_GRANT_TYPE = "grant_type"
)

func WxLogin(code string, channel int32) (openID string, sessionKey string,error error) {
	var url string
	if channel == constant.CHANNEL_TOUTIAO {
		url = "https://developer.toutiao.com/api/apps/jscode2session?appid=" + config.Server.TtConfig.AppID +
			"&secret=" + config.Server.TtConfig.AppSecret +
			"&code=" + code
	} else if channel== constant.CHANNEL_WECHAT {
		url = "https://api.weixin.qq.com/sns/jscode2session?appid=" + config.Server.WxConfig.AppID +
			"&secret=" + config.Server.WxConfig.AppSecret +
			"&grant_type=authorization_code" +
			"&js_code=" + code
	} else {
		log.Fatal("platform config is nil")
	}

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var openid string
	var session_key string
	if bytes.Contains(body, []byte("openid")) {
		atr := &LoginResponse{}
		err = json.Unmarshal(body, &atr)
		openid = atr.OpenID
		session_key = atr.SessionKey
		if err != nil {
			return "", "", err
		}
	} else {
		ater := &ErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			return "", "", err
		}
		return "", "", fmt.Errorf("%s", ater.ErrMsg)
	}
	return openid, session_key, nil
}

func WxDecrypt(encryptedData, ivData string, session_key string) {
	decryptWechatAppletUser(encryptedData, ivData, session_key)
}

