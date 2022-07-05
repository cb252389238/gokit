package ep

import (
	"fmt"
	"log"
	"regexp"
	"unicode"
)

//正则判断是否是邮箱
func IsEmail(str string) bool {
	result, err := regexp.MatchString(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, str)
	if err != nil {
		log.Fatal(err)
	}
	if result {
		return true
	} else {
		return false
	}
}

//正则判断是否是网址
func IsUrl(str string) bool {
	result, err := regexp.MatchString(`^[a-zA-z]+://(\w+(-\w+)*)(\.(\w+(-\w+)*))*(\?\S*)?$`, str)
	if err != nil {
		log.Fatal(err)
	}
	if result {
		return true
	} else {
		return false
	}
}

//正则判断是否是网址
func IsMobile(str string) bool {
	result, err := regexp.MatchString(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$`, str)
	if err != nil {
		log.Fatal(err)
	}
	if result {
		return true
	} else {
		return false
	}
}

var zimu = &unicode.RangeTable{
	R16: []unicode.Range16{
		{
			Lo:     65,
			Hi:     90,
			Stride: 1,
		},
		{
			Lo:     97,
			Hi:     122,
			Stride: 1,
		},
	},
	R32: []unicode.Range32{
		{
			Lo:     65,
			Hi:     90,
			Stride: 1,
		},
		{
			Lo:     97,
			Hi:     122,
			Stride: 1,
		},
	},
}

var shuzi = &unicode.RangeTable{
	R16: []unicode.Range16{
		{
			Lo:     48,
			Hi:     57,
			Stride: 1,
		},
	},
	R32: []unicode.Range32{
		{
			Lo:     48,
			Hi:     57,
			Stride: 1,
		},
	},
}

var hanzi = &unicode.RangeTable{
	R16: []unicode.Range16{
		//基础汉字
		{
			Lo:     19968,
			Hi:     40869,
			Stride: 1,
		},
	},
	R32: []unicode.Range32{
		{
			Lo:     19968,
			Hi:     40869,
			Stride: 1,
		},
	},
}

var tables = []*unicode.RangeTable{
	unicode.Space,
	zimu,
	hanzi,
	shuzi,
	unicode.Punct,
}

func IsChinese(str string) bool {
	for _, r := range str {
		if unicode.IsOneOf(tables, r) {
			fmt.Printf("%s", string(r))
			return true
		}
	}
	return false
}
