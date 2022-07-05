package encrypt

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha2(str, param string) string {
	var result string
	switch param {
	case "256":
		result = sha2_256(str)
	case "512":
		result = sha2_512(str)
	default:
		result = ""
	}
	return result
}

func sha2_256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func sha2_512(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
