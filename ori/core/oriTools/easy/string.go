package easy

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"hash/crc32"
	"html"
	"io"
	"io/ioutil"
	"math/rand"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

// 字符串长度
func MbStrlen(str string) int {
	return utf8.RuneCountInString(str)
}

// 字符串反转
func StrRev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 随机打乱字符串
func StrShuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

// 解析字符串
func ParseStr(encodedString string, result map[string]any) error {
	var build func(map[string]any, []string, any) error
	build = func(result map[string]any, keys []string, value any) error {
		length := len(keys)
		key := strings.Trim(keys[0], "'\"")
		if length == 1 {
			result[key] = value
			return nil
		}
		if keys[1] == "" && length == 2 {
			if key == "" {
				return nil
			}
			val, ok := result[key]
			if !ok {
				result[key] = []any{value}
				return nil
			}
			children, ok := val.([]any)
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			result[key] = append(children, value)
			return nil
		}

		if keys[1] == "" && length > 2 && keys[2] != "" {
			val, ok := result[key]
			if !ok {
				result[key] = []any{}
				val = result[key]
			}
			children, ok := val.([]any)
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			if l := len(children); l > 0 {
				if child, ok := children[l-1].(map[string]any); ok {
					if _, ok := child[keys[2]]; !ok {
						_ = build(child, keys[2:], value)
						return nil
					}
				}
			}
			child := map[string]any{}
			_ = build(child, keys[2:], value)
			result[key] = append(children, child)

			return nil
		}

		val, ok := result[key]
		if !ok {
			result[key] = map[string]any{}
			val = result[key]
		}
		children, ok := val.(map[string]any)
		if !ok {
			return fmt.Errorf("expected type 'map[string]interface{}' for key '%s', but got '%T'", key, val)
		}

		return build(children, keys[1:], value)
	}

	parts := strings.Split(encodedString, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			return err
		}
		for key[0] == ' ' {
			key = key[1:]
		}
		if key == "" || key[0] == '[' {
			continue
		}
		value, err := url.QueryUnescape(part[pos+1:])
		if err != nil {
			return err
		}

		var keys []string
		left := 0
		for i, k := range key {
			if k == '[' && left == 0 {
				left = i
			} else if k == ']' {
				if left > 0 {
					if len(keys) == 0 {
						keys = append(keys, key[:left])
					}
					keys = append(keys, key[left+1:i])
					left = 0
					if i+1 < len(key) && key[i+1] != '[' {
						break
					}
				}
			}
		}
		if len(keys) == 0 {
			keys = append(keys, key)
		}
		first := ""
		for i, chr := range keys[0] {
			if chr == ' ' || chr == '.' || chr == '[' {
				first += "_"
			} else {
				first += string(chr)
			}
			if chr == '[' {
				first += keys[0][i+1:]
				break
			}
		}
		keys[0] = first
		if err := build(result, keys, value); err != nil {
			return err
		}
	}
	return nil
}

// crc32
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func UrlEncode(str string) string {
	return url.QueryEscape(str)
}

func UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// html解码
func HtmlDecode(str string) string {
	return html.UnescapeString(str)
}

// html编码
func HtmlEncode(str string) string {
	return html.EscapeString(str)
}

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// 转换字符编码
// dstCharset 目标编码   srcCharset 源编码
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := io.ReadAll(
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

// 字符转utf8
// srcCharset 源编码 src 需要转的字符串
func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return Convert("UTF-8", srcCharset, src)
}

// utf8转其他编码
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
