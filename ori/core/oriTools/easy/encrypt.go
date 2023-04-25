package easy

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/sha3"
	"strings"
	"time"
)

// md5加密
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

// 3des加密
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

// 3des 解密
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

// aes加密
func AesEncrypt(orig, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pkcs7Padding_aes(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

// aes解密
func AesDecrypt(cryted, key string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = pkcs7UnPadding_aes(orig)
	return string(orig)
}

// 补码
func pkcs7Padding_aes(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func pkcs7UnPadding_aes(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// des加密
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

type Rsa struct {
	privateKey    string
	publicKey     string
	rsaPrivateKey *rsa.PrivateKey
	rsaPublicKey  *rsa.PublicKey
}

func NewRsa(publicKey, privateKey string) *Rsa {
	rsaObj := &Rsa{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	rsaObj.init() //初始化，如果存在公钥私钥，将其解析
	return rsaObj
}

// 初始化
func (r *Rsa) init() {
	if r.privateKey != "" {
		//将私钥解码
		block, _ := pem.Decode([]byte(r.privateKey))
		//pkcs1   //判断是否包含 BEGIN RSA 字符串,这个是由下面生成的时候定义的
		if strings.Index(r.privateKey, "BEGIN RSA") > 0 {
			//解析私钥
			r.rsaPrivateKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
		} else { //pkcs8
			//解析私钥
			privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
			//转换格式  类型断言
			r.rsaPrivateKey = privateKey.(*rsa.PrivateKey)
		}
	}
	if r.publicKey != "" {
		//将公钥解码 解析 转换格式
		block, _ := pem.Decode([]byte(r.publicKey))
		publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
		r.rsaPublicKey = publicKey.(*rsa.PublicKey)
	}
}

// Encrypt 加密
func (r *Rsa) Encrypt(data []byte) ([]byte, error) {
	// blockLength = 密钥长度 = 一次能加密的明文长度
	// "/8" 将bit转为bytes
	// "-11" 为 PKCS#1 建议的 padding 占用了 11 个字节
	blockLength := r.rsaPublicKey.N.BitLen()/8 - 11
	//如果明文长度不大于密钥长度，可以直接加密
	if len(data) <= blockLength {
		//对明文进行加密
		return rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, []byte(data))
	}
	//否则分段加密
	//创建一个新的缓冲区
	buffer := bytes.NewBufferString("")
	pages := len(data) / blockLength //切分为多少块
	//循环加密
	for i := 0; i <= pages; i++ {
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages { //最后一页的判断
			if start == len(data) {
				continue
			}
			end = len(data)
		}
		//分段加密
		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, r.rsaPublicKey, data[start:end])
		if err != nil {
			return nil, err
		}
		//写入缓冲区
		buffer.Write(chunk)
	}
	//读取缓冲区内容并返回，即返回加密结果
	return buffer.Bytes(), nil
}

// Decrypt 解密
func (r *Rsa) Decrypt(data []byte) ([]byte, error) {
	//加密后的密文长度=密钥长度。如果密文长度大于密钥长度，说明密文非一次加密形成
	//1、获取密钥长度
	blockLength := r.rsaPublicKey.N.BitLen() / 8
	if len(data) <= blockLength { //一次形成的密文直接解密
		return rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, data)
	}

	buffer := bytes.NewBufferString("")
	pages := len(data) / blockLength
	for i := 0; i <= pages; i++ { //循环解密
		start := i * blockLength
		end := (i + 1) * blockLength
		if i == pages {
			if start == len(data) {
				continue
			}
			end = len(data)
		}
		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, r.rsaPrivateKey, data[start:end])
		if err != nil {
			return nil, err
		}
		buffer.Write(chunk)
	}
	return buffer.Bytes(), nil
}

// Sign 签名
func (r *Rsa) Sign(data []byte, sHash crypto.Hash) ([]byte, error) {
	hash := sHash.New()
	hash.Write(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, r.rsaPrivateKey, sHash, hash.Sum(nil))
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// Verify 验签
func (r *Rsa) Verify(data []byte, sign []byte, sHash crypto.Hash) bool {
	h := sHash.New()
	h.Write(data)
	return rsa.VerifyPKCS1v15(r.rsaPublicKey, sHash, h.Sum(nil), sign) == nil
}

// CreateKeys 生成pkcs1 格式的公钥私钥
func (r *Rsa) CreateKeys(keyLength int) (privateKey, publicKey string) {
	//根据 随机源 与 指定位数，生成密钥对。rand.Reader = 密码强大的伪随机生成器的全球共享实例
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}
	//编码私钥
	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY", //自定义类型
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPrivateKey),
	}))
	//编码公钥
	objPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}
	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: objPkix,
	}))
	return
}

// CreatePkcs8Keys 生成pkcs8 格式公钥私钥
func (r *Rsa) CreatePkcs8Keys(keyLength int) (privateKey, publicKey string) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
	if err != nil {
		return
	}
	//两种方式
	//一：1、生成pkcs1格式的密钥 2、将其转化为pkcs8格式的密钥（使用自定义方法）
	//  objPkcs1 := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	//  objPkcs8 := r.Pkcs1ToPkcs8(objPkcs1)
	//二：直接使用 x509 包 MarshalPKCS8PrivateKey 生成pkcs8密钥
	objPkcs8, _ := x509.MarshalPKCS8PrivateKey(rsaPrivateKey)
	privateKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: objPkcs8,
	}))
	objPkix, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)
	if err != nil {
		return
	}
	publicKey = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: objPkix,
	}))
	return
}

// Pkcs1ToPkcs8 将pkcs1 转到 pkcs8 自定义
func (r *Rsa) Pkcs1ToPkcs8(key []byte) []byte {
	info := struct {
		Version             int
		PrivateKeyAlgorithm []asn1.ObjectIdentifier
		PrivateKey          []byte
	}{}
	info.Version = 0
	info.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 1)
	info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}
	info.PrivateKey = key
	k, _ := asn1.Marshal(info)
	return k
}

func Sha1(str string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(str))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// binary true 返回二进制
func HmacSHA1(key string, data string, binary bool) interface{} {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	if !binary {
		return hex.EncodeToString(mac.Sum(nil))
	} else {
		return mac.Sum(nil)
	}
}

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
