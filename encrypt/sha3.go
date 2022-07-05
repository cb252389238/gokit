package encrypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
	"time"
)

func Sha3(str, param string) string {
	var result string
	switch param {
	case "512":
		result = sha3_512(str)
	case "384":
		result = sha3_384(str)
	case "256":
		result = sha3_256(str)
	case "224":
		result = sha3_224(str)
	default:
		result = ""
	}
	return result
}

func sha3_512(str string) string {
	h := sha3.New512()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha3_384(str string) string {
	h := sha3.New384()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha3_256(str string) string {
	h := sha3.New256()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha3_224(str string) string {
	h := sha3.New224()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptHmacSHA256(ak, sk string) (sign, ts string) {
	ts = time.Now().UTC().Format("20060102150405000")
	h := hmac.New(sha256.New, []byte(sk))
	h.Write([]byte(ts + ak))
	sign = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}
