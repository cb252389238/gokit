package simhash2

import (
	"fmt"
	"testing"
)

func TestCalculateTFIDF(t *testing.T) {
	s1 := []string{"今天", "我", "参加", "同学", "婚礼", "同学", "婚礼", "举行", "非常", "盛大", "祝愿", "同学", "新婚", "快乐", "百年好合"}
	tfidf := calculateTFIDF(s1, [][]string{{"我", "婚礼"}})
	fmt.Println(tfidf)
}

func TestSimhash(t *testing.T) {
	s1 := []string{"今天", "我", "参加", "同学", "婚礼", "同学", "婚礼", "举行", "非常", "盛大", "祝愿", "同学", "新婚", "快乐", "百年好合"}
	Simhash(s1)
}
