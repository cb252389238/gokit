package ep

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"sync"
)

var (
	StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
)

/**
获取指定长度的随机数字字符串
*/
func RandNumStr(len int) string {
	var minNum int64 = 1
	for i := 1; i < len; i++ {
		minNum *= 10
	}
	var maxNum int64 = 1
	for j := 0; j < len; j++ {
		maxNum *= 10
	}
	maxNum = maxNum - 1
	maxBigInt := big.NewInt(maxNum)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	var randNum string
	if i.Int64() < minNum {
		return RandNumStr(len)
	} else {
		randNum = strconv.FormatInt(i.Int64(), 10)
		return randNum
	}

}

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

type Random struct {
	L        sync.Mutex
	RandSeek int64
}

/*
	获取随机数
*/
func RandNum(max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64()
}

func GetChNum(len int) string {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			case ch <- 2:
			case ch <- 3:
			case ch <- 4:
			case ch <- 5:
			case ch <- 6:
			case ch <- 7:
			case ch <- 8:
			case ch <- 9:
			}
		}
	}()
	var str string
	for i := 0; i < len; i++ {
		if v, ok := <-ch; ok {
			str += strconv.Itoa(v)
		}
	}
	if []byte(str)[0] == 48 {
		str = str[1:] + "0"
	}
	return str
}
