package charchar

import (
	"strings"
	"unicode"
)

// 去除标点符号
func RemovePunct(s string) string {
	var result []rune
	for _, r := range s {
		if !unicode.IsPunct(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

// 去除无意义词
var nonsenseWords = map[string]struct{}{
	"的": {},
	"了": {},
	"在": {},
	"于": {},
	"又": {},
	"我": {},
	"很": {},
}

func RemoveNonsenseWord(s string) string {
	newWords := []rune{}
	for _, word := range s {
		if _, ok := nonsenseWords[string(word)]; !ok {
			newWords = append(newWords, word)
		}
	}
	return string(newWords)
}

func GetTF(s string) map[string]int {
	frequencies := make(map[string]int)
	for _, v := range s {
		word := strings.TrimSpace(string(v))
		if len(word) > 0 {
			frequencies[word]++
		}
	}

	return frequencies
}
