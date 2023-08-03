package minhash

import (
	"fmt"
	"testing"
)

func TestComputeSimilarity(t *testing.T) {
	s1 := []string{"今天", "我", "参加", "同学", "婚礼", "同学", "婚礼", "举行", "非常", "盛大", "祝愿", "同学", "新婚", "快乐", "百年好合"}
	s2 := []string{"今天", "我", "参加", "同学", "同学", "婚礼", "举行", "非常", "盛大", "祝愿", "新婚", "快乐", "百年好合"}
	h1 := ComputeMinHashSignature(s1)
	h2 := ComputeMinHashSignature(s2)
	similarity := ComputeSimilarity(h1, h2)
	fmt.Println(similarity)
}
