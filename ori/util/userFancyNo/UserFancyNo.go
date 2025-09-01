package core

import (
	"errors"
	"strconv"
)

// 用户靓号服务
type UserFancyNo struct {
}

var (
	fancyNoTypesMapLevel = map[string]int{
		"AAAA":     5,
		"AAAAA":    5,
		"AAAAAA":   5,
		"AAAAAAA":  5,
		"AAAAAAAA": 5,
		"ABABAB":   4,
		"ABBBABBB": 4,
		"BBBABBBA": 4,
		"ABCDEFG":  4,
		"GFEDCBA":  4,
		"ABCDEAAA": 4,
		"EDCBAAAA": 4,
		"ABCDEFAA": 4,
		"FEDCBAAA": 4,
		"ABCDEABC": 4,
		"ABCDECBA": 4,
		"EDCBAABC": 3,
		"EDCBACBA": 3,
		"ABCDEFAB": 4,
		"ABCDEFBA": 4,
		"FEDCBAAB": 3,
		"FEDCBABA": 3,
		"AAAABCDE": 4,
		"AAAEDCBA": 4,
		"AAABCDEF": 4,
		"AAFEDCBA": 3,
		"ABCABCDE": 4,
		"ABCEDCBA": 4,
		"CBAABCDE": 4,
		"CBAEDCBA": 3,
		"ABABCDEF": 4,
		"ABFEDCBA": 3,
		"BAABCDEF": 4,
		"BAFEDCBA": 3,
	}
)

func (u *UserFancyNo) GetFancyNoLevel(types string) int {
	return fancyNoTypesMapLevel[types]
}

func (u *UserFancyNo) GetFancyNoTypesByLevel(level int) (res []string) {
	res = []string{}
	for k, v := range fancyNoTypesMapLevel {
		if v == level {
			res = append(res, k)
		}
	}
	return
}

// 判断是否为靓号
func (u *UserFancyNo) IsFancyNo(idStr string) (bool, string, error) {
	// 规则0：校验是否为8位数字
	if len(idStr) != 8 {
		return false, "", errors.New("invalid length")
	}
	if id, err := strconv.Atoi(idStr); err != nil || id < 10000000 || id > 99999999 {
		return false, "", errors.New("invalid range")
	}

	ok, types := u.checkRulesConcurrently(idStr)
	return ok, types, nil
}

type checkRes struct {
	IsFancyNo bool
	Types     string
}

// 返回所有靓号类型
func (u *UserFancyNo) GetFancyNoTypes() (types []string) {
	types = []string{
		"AAAA",
		"AAAAA",
		"AAAAAA",
		"AAAAAAA",
		"AAAAAAAA",
		"ABABAB",
		"ABBBABBB",
		"BBBABBBA",
		"ABCDEFG",
		"GFEDCBA",
		"ABCDEAAA",
		"EDCBAAAA",
		"ABCDEFAA",
		"FEDCBAAA",
		"ABCDEABC",
		"ABCDECBA",
		"EDCBAABC",
		"EDCBACBA",
		"ABCDEFAB",
		"ABCDEFBA",
		"FEDCBAAB",
		"FEDCBABA",
		"AAAABCDE",
		"AAAEDCBA",
		"AAABCDEF",
		"AAFEDCBA",
		"ABCABCDE",
		"ABCEDCBA",
		"CBAABCDE",
		"CBAEDCBA",
		"ABABCDEF",
		"ABFEDCBA",
		"BAABCDEF",
		"BAFEDCBA",
	}
	return
}

// 检查所有规则
func (u *UserFancyNo) checkRulesConcurrently(s string) (bool, string) {
	checkers := []func(string) (bool, string){
		u.checkConsecutive,
		u.checkABABAB,
		u.checkABBBABBB,
		u.checkBBBABBBA,
		u.checkAscending,
		u.checkFront5IncDecTail3Same,
		u.checkFront6IncDecTail2Same,
		u.checkFront5IncDecTail3IncDec,
		u.checkFront6IncDecTail2IncDec,
		u.checkFront2SameTail6IncDec,
		u.checkFront3IncDecTail5IncDec,
		u.checkFront2IncDecTail6IncDec,
	}

	results := make(chan checkRes, len(checkers))
	for _, f := range checkers {
		go func(fn func(string) (bool, string)) {
			ok, types := fn(s)
			results <- checkRes{
				IsFancyNo: ok,
				Types:     types,
			}
		}(f)
	}
	for i := 0; i < len(checkers); i++ {
		v := <-results
		if v.IsFancyNo {
			return v.IsFancyNo, v.Types
		}
	}
	return false, ""
}

/* ------------------- 规则实现 ------------------- */

// 规则1：连续出现4-8次相同数字
func (u *UserFancyNo) checkConsecutive(s string) (bool, string) {
	maxCount := 1
	current := rune(s[0])
	count := 1
	for _, c := range s[1:] {
		if c == current {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			current = c
			count = 1
		}
	}
	switch {
	case maxCount == 4:
		return true, "AAAA"
	case maxCount == 5:
		return true, "AAAAA"
	case maxCount == 6:
		return true, "AAAAAA"
	case maxCount == 7:
		return true, "AAAAAAA"
	case maxCount == 8:
		return true, "AAAAAAAA"
	}
	return false, ""
}

// 规则2：ababab模式 (如23232323)
func (u *UserFancyNo) checkABABAB(s string) (bool, string) {
	// 遍历所有可能的六位窗口
	for i := 0; i <= 2; i++ { // 检查起始位置0,1,2的六位
		a, b := s[i], s[i+1]
		if a == b {
			continue // 排除a=b的情况
		}
		// 检查后续位是否满足ababab
		if s[i+2] == a && s[i+3] == b &&
			s[i+4] == a && s[i+5] == b {
			return true, "ABABAB"
		}
	}
	return false, ""
}

// 规则3：abbbabbb模式 (如12221222)
func (u *UserFancyNo) checkABBBABBB(s string) (bool, string) {
	if s[0] == s[4] && s[1] == s[2] && s[2] == s[3] &&
		s[5] == s[6] && s[6] == s[7] && s[1] == s[5] && s[0] != s[1] {
		return true, "ABBBABBB"
	} else {
		return false, ""
	}
}

// 规则3：bbbabbba模式
func (u *UserFancyNo) checkBBBABBBA(s string) (bool, string) {
	if s[3] == s[7] && s[0] == s[1] && s[1] == s[2] &&
		s[4] == s[5] && s[5] == s[6] && s[0] == s[4] && s[0] != s[3] {
		return true, "BBBABBBA"
	} else {
		return false, ""
	}
}

// 规则4：连续7递增,递减
func (u *UserFancyNo) checkAscending(s string) (bool, string) {
	// 检查所有可能的7位窗口
	for i := 0; i <= 1; i++ { // 起始位置0和1
		ascValid, descValid := true, true
		// 遍历当前窗口的6次数字变化
		for j := i; j < i+6; j++ {
			// 检查递增
			if s[j+1]-s[j] != 1 {
				ascValid = false
			}
			// 检查递减
			if s[j]-s[j+1] != 1 {
				descValid = false
			}
			// 提前退出优化
			if !ascValid && !descValid {
				break
			}
		}
		// 任意模式匹配则返回true
		if ascValid {
			return true, "ABCDEFG"
		}
		if descValid {
			return true, "GFEDCBA"
		}
	}
	return false, ""
}

// 规则5: 前5递增/递减 + 后3相同 (如12345555)
func (u *UserFancyNo) checkFront5IncDecTail3Same(s string) (bool, string) {
	front := s[:5]
	tail := s[5:]
	if u.isIncreasing(front) && u.allSame(tail) {
		return true, "ABCDEAAA"
	}
	if u.isDecreasing(front) && u.allSame(tail) {
		return true, "EDCBAAAA"
	}
	return false, ""
}

// 规则6: 前6递增/递减 + 后2相同 (如12345677)
func (u *UserFancyNo) checkFront6IncDecTail2Same(s string) (bool, string) {
	front := s[:6]
	tail := s[6:]
	if u.isIncreasing(front) && u.allSame(tail) {
		return true, "ABCDEFAA"
	}
	if u.isDecreasing(front) && u.allSame(tail) {
		return true, "FEDCBAAA"
	}
	return false, ""
}

// 规则7: 前5递增/递减 + 后3递增/递减 (如12345456)
func (u *UserFancyNo) checkFront5IncDecTail3IncDec(s string) (bool, string) {
	front := s[:5]
	tail := s[5:]
	if u.isIncreasing(front) && u.isIncreasing(tail) {
		return true, "ABCDEABC"
	}
	if u.isIncreasing(front) && u.isDecreasing(tail) {
		return true, "ABCDECBA"
	}
	if u.isDecreasing(front) && u.isIncreasing(tail) {
		return true, "EDCBAABC"
	}
	if u.isDecreasing(front) && u.isDecreasing(tail) {
		return true, "EDCBACBA"
	}
	return false, ""
}

// 规则8: 前6递增/递减 + 后2递增/递减 (如12345678)
func (u *UserFancyNo) checkFront6IncDecTail2IncDec(s string) (bool, string) {
	front := s[:6]
	tail := s[6:]
	if u.isIncreasing(front) && u.isIncreasing(tail) {
		return true, "ABCDEFAB"
	}
	if u.isIncreasing(front) && u.isDecreasing(tail) {
		return true, "ABCDEFBA"
	}
	if u.isDecreasing(front) && u.isIncreasing(tail) {
		return true, "FEDCBAAB"
	}
	if u.isDecreasing(front) && u.isDecreasing(tail) {
		return true, "FEDCBABA"
	}
	return false, ""
}

// 规则9: 前3相同 + 后5递增/递减 (如55512345)
func (u *UserFancyNo) checkFront3SameTail5IncDec(s string) (bool, string) {
	front := s[:3]
	tail := s[3:]
	if u.allSame(front) && u.isIncreasing(tail) {
		return true, "AAAABCDE"
	}
	if u.allSame(front) && u.isDecreasing(tail) {
		return true, "AAAEDCBA"
	}
	return false, ""
}

// 规则10: 前2相同 + 后6递增/递减 (如33123456)
func (u *UserFancyNo) checkFront2SameTail6IncDec(s string) (bool, string) {
	front := s[:2]
	tail := s[2:]
	if u.allSame(front) && u.isIncreasing(tail) {
		return true, "AAABCDEF"
	}
	if u.allSame(front) && u.isDecreasing(tail) {
		return true, "AAFEDCBA"
	}
	return false, ""
}

// 规则11: 前3递增/递减 + 后5递增/递减 (如12345678)
func (u *UserFancyNo) checkFront3IncDecTail5IncDec(s string) (bool, string) {
	front := s[:3]
	tail := s[3:]
	if u.isIncreasing(front) && u.isIncreasing(tail) {
		return true, "ABCABCDE"
	}
	if u.isIncreasing(front) && u.isDecreasing(tail) {
		return true, "ABCEDCBA"
	}
	if u.isDecreasing(front) && u.isIncreasing(tail) {
		return true, "CBAABCDE"
	}
	if u.isDecreasing(front) && u.isDecreasing(tail) {
		return true, "CBAEDCBA"
	}
	return false, ""
}

// 规则12: 前2递增/递减 + 后6递增/递减 (如12345678)
func (u *UserFancyNo) checkFront2IncDecTail6IncDec(s string) (bool, string) {
	front := s[:2]
	tail := s[2:]
	if u.isIncreasing(front) && u.isIncreasing(tail) {
		return true, "ABABCDEF"
	}
	if u.isIncreasing(front) && u.isDecreasing(tail) {
		return true, "ABFEDCBA"
	}
	if u.isDecreasing(front) && u.isIncreasing(tail) {
		return true, "BAABCDEF"
	}
	if u.isDecreasing(front) && u.isDecreasing(tail) {
		return true, "BAFEDCBA"
	}
	return false, ""
}

// 检查是否严格递增
func (u *UserFancyNo) isIncreasing(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i+1]-s[i] != 1 {
			return false
		}
	}
	return len(s) > 0
}

// 检查是否递减
func (u *UserFancyNo) isDecreasing(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i]-s[i+1] != 1 {
			return false
		}
	}
	return len(s) > 0
}

// 检查是否全相同
func (u *UserFancyNo) allSame(s string) bool {
	if len(s) == 0 {
		return true
	}
	first := s[0]
	for _, c := range s[1:] {
		if byte(c) != first {
			return false
		}
	}
	return true
}
