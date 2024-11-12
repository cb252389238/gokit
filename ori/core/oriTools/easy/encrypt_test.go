package easy

import (
	"crypto"
	"testing"
)

var (
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBAKOFqluzvKBxVWdYi2fNQCYPWJ7sC0mqH0WLrsnCTmSWDZMhm6pp
IMxzV6OElphDGi/hxgQBg9dwn8xDkqJzqwcCAwEAAQJBAIwZGATrMB+yGf6qEP4F
DwHMwhuelmktlQ9Lhpwbmnh37Dw8Vp5yKAEqpTtmYD0KG7QzlU5cYaTtFioS6A2l
OiECIQDSerkA2Doi3FbjlrhIsWR88EneyS2NUaq+twh2VeZn9wIhAMbjIpobMCS1
b0C+d3Bum399C/1hDfdiAtUCWbwZkrFxAiBg64X5L5hFqTSRhDvDrXvaVEOPxQ+m
vW5kd5/77b41LQIhAJgLCefPwxU9EsjnEr4EAKIMwX65lIi7B7k5q8oNrsQxAiEA
0Tx3+3BFtT5cFZroNn85W8iqxd4p82Dh31IuRasNvt4=
-----END RSA PRIVATE KEY-----`
	publicKey = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAKOFqluzvKBxVWdYi2fNQCYPWJ7sC0mq
H0WLrsnCTmSWDZMhm6ppIMxzV6OElphDGi/hxgQBg9dwn8xDkqJzqwcCAwEAAQ==
-----END PUBLIC KEY-----`
)

func TestMd5(t *testing.T) {
	t.Log(Md5("123456", 32, true))
	t.Log(Md5("123456", 16, false))
}

func TestTripleDesEncrypt(t *testing.T) {
	key := "akendig1241258ds3d4wjsid"
	encrypt, err := TripleDesEncrypt("123456", key)
	if err != nil {
		t.Error(err)
	}
	t.Log(encrypt)
	decrypt, err := TripleDesDecrypt(encrypt, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(decrypt)
}

func TestAesEncrypt(t *testing.T) {
	key := "akendig1241258ds3d4wjsid"
	encrypt := AesEncrypt("123456", key)
	t.Log(encrypt)
	decrypt := AesDecrypt(encrypt, key)
	t.Log(decrypt)
}

func TestDesEncrypt(t *testing.T) {
	key := "12345678"
	encrypt, err := DesEncrypt("123456", key)
	if err != nil {
		t.Error(err)
	}
	t.Log(encrypt)
	decrypt, err := DesDecrypt(encrypt, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(decrypt)
}

func TestRsa(t *testing.T) {
	rsa := NewRsa(publicKey, privateKey)
	encrypt, err := rsa.Encrypt([]byte("123456"))
	if err != nil {
		t.Error(err)
	}
	t.Log(string(encrypt))
	decrypt, err := rsa.Decrypt(encrypt)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(decrypt))
	sign, err := rsa.Sign([]byte("123456"), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(sign))
	verify := rsa.Verify([]byte("123456"), sign, crypto.SHA256)
	t.Log(verify)
	key, k, err := rsa.CreateKeys(50)
	t.Log(key, k, err)
}

func TestSha1(t *testing.T) {
	t.Log(Sha1("123456"))
	t.Log(Sha2("123456", "256"))
	t.Log(Sha2("123456", "512"))
	t.Log(Sha3("123456", "256"))
	t.Log(Sha3("123456", "384"))
	t.Log(Sha3("123456", "256"))
	t.Log(Sha3("123456", "224"))
}

func TestHmacSHA1(t *testing.T) {
	t.Log(HmacSHA1("ak", "123456"))
}
