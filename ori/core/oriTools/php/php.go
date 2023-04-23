package php

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Empty(val interface{}) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func Is_numeric(val any) bool {
	switch val := val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

func Exit(status int) {
	os.Exit(status)
}

func Die(status int) {
	os.Exit(status)
}

func Getenv(varname string) string {
	return os.Getenv(varname)
}

func Putenv(setting string) error {
	s := strings.Split(setting, "=")
	if len(s) != 2 {
		panic("setting: invalid")
	}
	return os.Setenv(s[0], s[1])
}

/*
版本比较
*/
func Version_compare(version1, version2, operator string) bool {
	var vcompare func(string, string) int
	var canonicalize func(string) string
	var special func(string, string) int

	vcompare = func(origV1, origV2 string) int {
		if origV1 == "" || origV2 == "" {
			if origV1 == "" && origV2 == "" {
				return 0
			}
			if origV1 == "" {
				return -1
			}
			return 1
		}

		ver1, ver2, compare := "", "", 0
		if origV1[0] == '#' {
			ver1 = origV1
		} else {
			ver1 = canonicalize(origV1)
		}
		if origV2[0] == '#' {
			ver2 = origV2
		} else {
			ver2 = canonicalize(origV2)
		}
		n1, n2 := 0, 0
		for {
			p1, p2 := "", ""
			n1 = strings.IndexByte(ver1, '.')
			if n1 == -1 {
				p1, ver1 = ver1[:], ""
			} else {
				p1, ver1 = ver1[:n1], ver1[n1+1:]
			}
			n2 = strings.IndexByte(ver2, '.')
			if n2 == -1 {
				p2, ver2 = ver2, ""
			} else {
				p2, ver2 = ver2[:n2], ver2[n2+1:]
			}

			if (p1[0] >= '0' && p1[0] <= '9') && (p2[0] >= '0' && p2[0] <= '9') { // all is digit
				l1, _ := strconv.Atoi(p1)
				l2, _ := strconv.Atoi(p2)
				if l1 > l2 {
					compare = 1
				} else if l1 == l2 {
					compare = 0
				} else {
					compare = -1
				}
			} else if !(p1[0] >= '0' && p1[0] <= '9') && !(p2[0] >= '0' && p2[0] <= '9') { // all digit
				compare = special(p1, p2)
			} else {
				if p1[0] >= '0' && p1[0] <= '9' { // is digit
					compare = special("#N#", p2)
				} else {
					compare = special(p1, "#N#")
				}
			}

			if compare != 0 || n1 == -1 || n2 == -1 {
				break
			}
		}

		if compare == 0 {
			if ver1 != "" {
				if ver1[0] >= '0' && ver1[0] <= '9' {
					compare = 1
				} else {
					compare = vcompare(ver1, "#N#")
				}
			} else if ver2 != "" {
				if ver2[0] >= '0' && ver2[0] <= '9' {
					compare = -1
				} else {
					compare = vcompare("#N#", ver2)
				}
			}
		}

		return compare
	}

	canonicalize = func(version string) string {
		ver := []byte(version)
		l := len(ver)
		if l == 0 {
			return ""
		}
		var buf = make([]byte, l*2)
		j := 0
		for i, v := range ver {
			next := uint8(0)
			if i+1 < l { // Have the next one
				next = ver[i+1]
			}
			if v == '-' || v == '_' || v == '+' { // replace '-', '_', '+' to '.'
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else if (next > 0) &&
				(!(next >= '0' && next <= '9') && (v >= '0' && v <= '9')) ||
				(!(v >= '0' && v <= '9') && (next >= '0' && next <= '9')) {
				buf[j] = v
				j++
				if v != '.' && next != '.' {
					buf[j] = '.'
					j++
				}
				continue
			} else if !((v >= '0' && v <= '9') ||
				(v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z')) {
				if j > 0 && buf[j-1] != '.' {
					buf[j] = '.'
					j++
				}
			} else {
				buf[j] = v
				j++
			}
		}

		return string(buf[:j])
	}

	special = func(form1, form2 string) int {
		found1, found2, len1, len2 := -1, -1, len(form1), len(form2)
		forms := map[string]int{
			"dev":   0,
			"alpha": 1,
			"a":     1,
			"beta":  2,
			"b":     2,
			"RC":    3,
			"rc":    3,
			"#":     4,
			"pl":    5,
			"p":     5,
		}

		for name, order := range forms {
			if len1 < len(name) {
				continue
			}
			if strings.Compare(form1[:len(name)], name) == 0 {
				found1 = order
				break
			}
		}
		for name, order := range forms {
			if len2 < len(name) {
				continue
			}
			if strings.Compare(form2[:len(name)], name) == 0 {
				found2 = order
				break
			}
		}

		if found1 == found2 {
			return 0
		} else if found1 > found2 {
			return 1
		} else {
			return -1
		}
	}

	compare := vcompare(version1, version2)

	switch operator {
	case "<", "lt":
		return compare == -1
	case "<=", "le":
		return compare != 1
	case ">", "gt":
		return compare == 1
	case ">=", "ge":
		return compare != -1
	case "==", "=", "eq":
		return compare == 0
	case "!=", "<>", "ne":
		return compare != 0
	default:
		panic("operator: invalid")
	}
}
