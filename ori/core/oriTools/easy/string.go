package easy

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math/rand"
	"net/url"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

/*
返回字符串长度
*/
func Strlen(str string) int {
	return len(str)
}

/*
返回汉字长度
*/
func Mb_strlen(str string) int {
	return utf8.RuneCountInString(str)
}

/*
字符串替换
*/
func Str_replace(search, replace, subject string, count int) string {
	return strings.Replace(subject, search, replace, count)
}

/*
分割字符串
*/
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

/*
合并字符串
*/
func Implode(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

func Substr(str string, start uint, length int) string {
	if length < -1 {
		return str
	}
	switch {
	case length == -1:
		return str[start:]
	case length == 0:
		return ""
	}
	end := int(start) + length
	if end > len(str) {
		end = len(str)
	}
	return str[start:end]
}

/*
转成小写
*/
func Strtolower(str string) string {
	return strings.ToLower(str)
}

func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

/*
反转字符串
*/
func Strrev(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
字符重复n次
*/
func Str_repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}

/*
随机打乱字符串
*/
func Str_shuffle(str string) string {
	runes := []rune(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, len(runes))
	for i, v := range r.Perm(len(runes)) {
		s[i] = runes[v]
	}
	return string(s)
}

/*
解析字符串
*/
func Parse_str(encodedString string, result map[string]any) error {
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

/*
去除两边得某些字符串
*/
func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

func Ltrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	return strings.TrimLeft(str, characterMask[0])
}

func Rtrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	return strings.TrimRight(str, characterMask[0])
}

func Json_decode(data string, val any) error {
	return json.Unmarshal([]byte(data), val)
}

func Json_encode(val any) ([]byte, error) {
	return json.Marshal(val)
}

func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func Urlencode(str string) string {
	return url.QueryEscape(str)
}

func Urldecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

func Base64_encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64_decode(str string) (string, error) {
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

// 脱敏手机号
func PrivateMobile(mobile string) string {
	masked := ""
	if len(mobile) > 5 {
		masked = mobile[:3] + strings.Repeat("*", len(mobile)-5) + mobile[len(mobile)-2:]
	}
	return masked
}

// 脱敏名字
func PrivateRealName(name string) string {
	masked := ""
	runeName := []rune(name)
	switch {
	case len(runeName) > 3:
		masked = string(runeName[0]) + strings.Repeat("*", len(runeName)-2) + string(runeName[len(runeName)-1:])
	case len(runeName) == 3:
		masked = string(runeName[0]) + strings.Repeat("*", 1) + string(runeName[len(runeName)-1:])
	case len(runeName) == 2:
		masked = string(runeName[0]) + strings.Repeat("*", 1)
	default:
		masked = name
	}
	return masked
}

// 脱敏身份证号
func PrivateIdNo(idNo string) string {
	masked := ""
	switch {
	case len(idNo) > 14:
		masked = idNo[0:4] + strings.Repeat("*", len(idNo)-6) + idNo[len(idNo)-2:]
	case len(idNo) > 6 && len(idNo) <= 14:
		masked = idNo[0:3] + strings.Repeat("*", len(idNo)-5) + idNo[len(idNo)-2:]
	case len(idNo) <= 6 && len(idNo) > 3:
		masked = strings.Repeat("*", len(idNo)-2) + idNo[len(idNo)-2:]
	default:
		masked = strings.Repeat("*", len(idNo))
	}
	return masked
}
