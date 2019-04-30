package wx

import (

	"aliens/common/cipher"
	"encoding/base64"
	"encoding/json"
)

func decryptWechatAppletUser (encryptedData string, session_key string, ivData string) map[string]interface{} {
	encrypted, _ := base64.StdEncoding.DecodeString(encryptedData)
	aeskey, _ := base64.StdEncoding.DecodeString(session_key)
	iv, _ := base64.StdEncoding.DecodeString(ivData)

	decryptedData, _ := cipher.CBCIvDecrypt(encrypted, aeskey, iv)
	decryptedMapping := make(map[string]interface{})
	json.Unmarshal(decryptedData, &decryptedMapping)
	return decryptedMapping
}