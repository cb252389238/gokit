package php

import (
	"fmt"
	"testing"
)

func TestStrlen(t *testing.T) {
	fmt.Println(Strlen("123哈哈"))
}

func TestMb_strlen(t *testing.T) {
	fmt.Println(Mb_strlen("123哈哈"))
}

func TestStr_replace(t *testing.T) {
	fmt.Println(Str_replace(",", "|", "哈哈,喜喜,呵呵", -1))
}

func TestExplode(t *testing.T) {
	str := "1,3,4,5"
	fmt.Println(Explode(",", str))
}

func TestImplode(t *testing.T) {
	s := []string{"1", "2", "3"}
	fmt.Println(Implode("-", s))
}

func TestSubstr(t *testing.T) {
	s := "01234"
	fmt.Println(Substr(s, 0, -1))
}

func TestStrtolower(t *testing.T) {
	fmt.Println(Strtolower("ABC"))
}

func TestStrtoupper(t *testing.T) {
	fmt.Println(Strtoupper("abc"))
}

func TestStrrev(t *testing.T) {
	fmt.Println(Strrev("abc"))
}

func TestStr_repeat(t *testing.T) {
	fmt.Println(Str_repeat("h", 10))
}

func TestStr_shuffle2(t *testing.T) {
	fmt.Println(Str_shuffle("12345678"))
}

func TestParse_str(t *testing.T) {
	str := "id=1&name=2&cc=3"
	result := map[string]any{}
	err := Parse_str(str, result)
	fmt.Println(result, err)
}

func TestTrim(t *testing.T) {
	str := " 123 "
	fmt.Println(Trim(str, " "))
}

func TestLtrim(t *testing.T) {
	str := " 123 "
	fmt.Println(Ltrim(str, " "))
}

func TestRtrim(t *testing.T) {
	str := " 123 "
	fmt.Println(Rtrim(str, " "))
}

func TestJson(t *testing.T) {
	m := map[string]any{
		"aa": "aa",
		"b":  1,
		"c":  3,
	}
	encode, err := Json_encode(m)
	fmt.Println(string(encode), err)
	a := struct {
		AA string `json:"aa"`
		B  int    `json:"b"`
		C  int    `json:"c"`
	}{}
	err = Json_decode(string(encode), &a)
	fmt.Println(a, err)
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123"))
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123"))
}

func TestCrc32(t *testing.T) {
	fmt.Println(Crc32("123"))
}

func TestUrl(t *testing.T) {
	url := "www.baidu.com?file=123"
	urlencode := Urlencode(url)
	fmt.Println(urlencode)
	fmt.Println(Urldecode(urlencode))
}

func TestBase64(t *testing.T) {
	str := "123456"
	encode := Base64_encode(str)
	fmt.Println(encode)
	fmt.Println(Base64_decode(encode))
}
