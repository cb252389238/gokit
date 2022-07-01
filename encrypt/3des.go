package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
)

func TripleDesEncrypt(orig, key string) (string, error) {
	if orig == "" {
		return "", errors.New("Encryption objects cannot be null")
	}
	if len(key) != 24 {
		return "", errors.New("The secret key length of 3DES must be 24 bits")
	}
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 3DES的秘钥长度必须为24位
	block, _ := des.NewTripleDESCipher(k)
	// 补全码
	origData = pkcs5Padding_3des(origData, block.BlockSize())
	// 设置加密方式
	blockMode := cipher.NewCBCEncrypter(block, k[:8])
	// 创建密文数组
	crypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

/**
 * 解密
 */
func TipleDesDecrypt(crypted, key string) (string, error) {
	if crypted == "" {
		return "", errors.New("ciphertext cannot be null")
	}
	if len(key) != 24 {
		return "", errors.New("The secret key length of 3DES must be 24 bits")
	}
	// 用base64转成字节数组
	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	// key转成字节数组
	k := []byte(key)
	block, _ := des.NewTripleDESCipher(k)
	blockMode := cipher.NewCBCDecrypter(block, k[:8])
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)
	origData = pkcs5UnPadding_3des(origData)
	return string(origData), nil
}

func pkcs5Padding_3des(orig []byte, size int) []byte {
	length := len(orig)
	padding := size - length%size
	paddintText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orig, paddintText...)
}

func pkcs5UnPadding_3des(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
