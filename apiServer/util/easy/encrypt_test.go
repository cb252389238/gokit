package easy

import (
	"fmt"
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

func TestGenerateRSAKey(t *testing.T) {
	rsa := NewRsa("", "")
	key, publicKey := rsa.CreateKeys(512)
	fmt.Println(key)
	fmt.Println(publicKey)
}

func TestRsa_Encrypt(t *testing.T) {
	rsa := NewRsa(publicKey, privateKey)
	encrypt, err := rsa.Encrypt([]byte("加密测试消息"))
	fmt.Println(encrypt, err)
	decrypt, err := rsa.Decrypt(encrypt)
	fmt.Println(string(decrypt), err)
}

func TestAesEncrypt(t *testing.T) {
	encrypt := AesEncrypt("123456", "6c6be19ba451d2fe921f5929a3b0b690")
	fmt.Println(encrypt)
}
