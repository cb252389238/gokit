package customEncrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"strings"
)

// Base62 编码/解码
const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(data []byte) string {
	encoded := make([]byte, 0, len(data)*2)
	for _, b := range data {
		val := int(b)
		quotient := val / 62
		remainder := val % 62
		encoded = append(encoded, base62Chars[quotient], base62Chars[remainder])
	}
	return string(encoded)
}

func Base62Decode(s string) ([]byte, error) {
	if len(s)%2 != 0 {
		return nil, errors.New("invalid base62 string length")
	}
	decoded := make([]byte, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		c1 := strings.IndexByte(base62Chars, s[i])
		c2 := strings.IndexByte(base62Chars, s[i+1])
		if c1 == -1 || c2 == -1 {
			return nil, errors.New("invalid base62 character")
		}
		val := c1*62 + c2
		if val > 255 {
			return nil, errors.New("invalid base62 value")
		}
		decoded[i/2] = byte(val)
	}
	return decoded, nil
}

// AES-CBC 加密
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 生成随机IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// PKCS#7填充
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)

	// 加密
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	// 返回IV + 密文
	return append(iv, ciphertext...), nil
}

// AES-CBC解密
func decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	// 分离IV和密文
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("invalid ciphertext length")
	}

	// 解密
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去除填充
	return pkcs7Unpad(ciphertext)
}

// PKCS#7填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data")
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, errors.New("invalid padding")
	}
	if len(data) < padding {
		return nil, errors.New("padding exceeds data length")
	}
	return data[:len(data)-padding], nil
}

// 加密
func EncryptAndEncode(plaintext string, key []byte) (string, error) {
	ciphertext, err := encrypt([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return Base62Encode(ciphertext), nil
}

// 解密
func DecodeAndDecrypt(encoded string, key []byte) (string, error) {
	ciphertext, err := Base62Decode(encoded)
	if err != nil {
		return "", err
	}
	plaintext, err := decrypt(ciphertext, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
