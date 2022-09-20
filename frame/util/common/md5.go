package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5(str string, len int, isUpper bool) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串
	md5Encode := hex.EncodeToString(h.Sum(nil))
	if len == 16 {
		md5Encode = md5Encode[8:24]
	}
	if isUpper {
		md5Encode = strings.ToUpper(md5Encode)
	}
	return md5Encode
}
