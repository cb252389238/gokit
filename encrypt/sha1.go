package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

//binary true 返回二进制
func HmacSHA1(key string, data string, binary bool) interface{} {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	if !binary {
		return hex.EncodeToString(mac.Sum(nil))
	} else {
		return mac.Sum(nil)
	}
}
