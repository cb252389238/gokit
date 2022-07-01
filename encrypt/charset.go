package encrypt

import (
	"bytes"
	"errors"
	"fmt"

	"io/ioutil"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// 检查是否支持
func Supported(charset string) bool {
	return getEncoding(charset) != nil
}

//转换字符编码
//dstCharset 目标编码   srcCharset 源编码
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", err
			}
			src = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupported srcCharset: %s", srcCharset))
		}
	}
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", errors.New(fmt.Sprintf("unsupported srcCharset: %s", dstCharset))
			}
			dst = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupported srcCharset: %s", dstCharset))
		}
	} else {
		dst = src
	}
	return dst, nil
}

//字符转utf8
//srcCharset 源编码 src 需要转的字符串
func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

//utf8转其他编码
func UTF8To(dstCharset string, src string) (dst string, err error) {
	return Convert(dstCharset, "UTF-8", src)
}

func getEncoding(charset string) encoding.Encoding {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	if e, err := ianaindex.MIB.Encoding(charset); err == nil && e != nil {
		return e
	}
	return nil
}
