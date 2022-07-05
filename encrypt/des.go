package encrypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"errors"
)

func DesEncrypt(orig, key string) (string, error) {
	if orig == "" {
		return "", errors.New("Encryption objects cannot be null")
	}
	if len(key) != 24 {
		return "", errors.New("The secret key length of DES must be 8 bits")
	}
	// 将加密内容和秘钥转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 秘钥分组
	block, _ := des.NewCipher(k)
	//将明文按秘钥的长度做补全操作
	origData = pkcs5Padding_des(origData, block.BlockSize())
	//设置加密方式－CBC
	blockMode := cipher.NewCBCDecrypter(block, k)
	//创建明文长度的字节数组
	crypted := make([]byte, len(origData))
	//加密明文
	blockMode.CryptBlocks(crypted, origData)
	//将字节数组转换成字符串，base64编码
	return base64.StdEncoding.EncodeToString(crypted), nil
}

/**
 * DES解密方法
 */
func DesDecrypt(data, key string) (string, error) {
	if data == "" {
		return "", errors.New("Encryption objects cannot be null")
	}
	if len(key) != 24 {
		return "", errors.New("The secret key length of DES must be 8 bits")
	}
	k := []byte(key)
	//将加密字符串用base64转换成字节数组
	crypted, _ := base64.StdEncoding.DecodeString(data)
	//将字节秘钥转换成block快
	block, _ := des.NewCipher(k)
	//设置解密方式－CBC
	blockMode := cipher.NewCBCEncrypter(block, k)
	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))
	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)
	//去掉加密时补全的部分
	origData = pkcs5UnPadding_des(origData)
	return string(origData), nil
}

/**
 * 实现明文的补全
 * 如果ciphertext的长度为blockSize的整数倍，则不需要补全
 * 否则差几个则被几个，例：差5个则补5个5
 */
func pkcs5Padding_des(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

/**
 * 实现去补码，PKCS5Padding的反函数
 */
func pkcs5UnPadding_des(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
