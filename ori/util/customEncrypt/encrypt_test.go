package customEncrypt

import "testing"

func TestEncryptAndEncode(t *testing.T) {
	key := make([]byte, 16)
	encryptAndEncode, _ := EncryptAndEncode("713466cafda549ac8c4161fc4c507924", key)
	t.Log(encryptAndEncode)
	decryptAndDecode, _ := DecodeAndDecrypt(encryptAndEncode, key)
	t.Log(decryptAndDecode)
}
