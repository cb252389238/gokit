package simhash

import (
	"github.com/cespare/xxhash/v2"
	"hash/fnv"
)

const (
	// SimhashBitSize 定义Simhash的位数
	SimhashBitSize = 64
)

// 获取hash值
func SimHash(tokens []string) uint64 {
	// 计算每个分词的权重
	weights := make(map[string]int)
	for _, token := range tokens {
		weights[token]++
	}

	// 创建特征向量
	features := make([]int, SimhashBitSize)

	// 根据权重计算特征向量
	for token, weight := range weights {
		hash := hash(token)
		for i := 0; i < SimhashBitSize; i++ {
			bit := (hash >> uint(i)) & 1
			if bit == 1 {
				features[i] += weight
			} else {
				features[i] -= weight
			}
		}
	}

	// 降维
	var simhash uint64
	for i := 0; i < SimhashBitSize; i++ {
		if features[i] > 0 {
			simhash |= (1 << uint(i))
		}
	}

	return simhash
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func xHash(s string) uint64 {
	sum64 := xxhash.Sum64([]byte(s))
	return sum64
}

// HammingDistance 计算两个Simhash值的汉明距离
func HammingDistance(hash1, hash2 uint64) int {
	xor := hash1 ^ hash2
	dist := 0
	for xor != 0 {
		dist++
		xor &= xor - 1
	}
	return dist
}

// Similarity 判断两个Simhash值的相似度（0到1之间的值）
func Similarity(hash1, hash2 uint64) (int, float64) {
	distance := HammingDistance(hash1, hash2)
	similarity := 1.0 - float64(distance)/SimhashBitSize
	return distance, similarity
}
