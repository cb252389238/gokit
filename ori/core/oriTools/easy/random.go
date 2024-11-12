package easy

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

var (
	StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

/*
获取随机的字符串
*/
func RandStr(len int, chars string) string {
	if chars != "" {
		StdChars = []byte(chars)
	}
	str := newLenChars(len, StdChars)
	return str
}

func newLenChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

/*
获取随机数
*/
func RandNum(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

/*
*
获取指定范围的随机数
*/
func N(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return intn(max-min+1) + min
	}
	return intn(max+(0-min)+1) - (0 - min)
}

func generateRandomData() []byte {
	data := make([]byte, 4)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(generateRandomData())) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}
