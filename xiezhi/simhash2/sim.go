package simhash2

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
)

const (
	// SimhashBitSize 定义Simhash的位数
	SimhashBitSize = 64
)

// 获取hash值
func Simhash(tokens []string) uint64 {
	// 计算每个分词的权重
	weights := make(map[string]int)
	for _, token := range tokens {
		weights[token]++
	}
	fmt.Println("分词权重:", weights)
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
	fmt.Println("向量特征:", features)
	// 降维
	var simhash uint64
	for i := 0; i < SimhashBitSize; i++ {
		if features[i] > 0 {
			simhash |= (1 << uint(i))
		}
	}
	binary := strconv.FormatInt(int64(simhash), 2)
	fmt.Println("二进制:", binary)
	fmt.Println("hash结果:", simhash)
	return simhash
}

// hash 计算字符串的哈希值
func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
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
func Similarity(hash1, hash2 uint64) float64 {
	distance := HammingDistance(hash1, hash2)
	fmt.Println("字符海明距离：", distance)
	similarity := 1.0 - float64(distance)/SimhashBitSize
	return similarity
}

// 计算词频tf
func calculateTermFrequency(words []string) map[string]float64 {
	tf := make(map[string]float64)
	totalWords := float64(len(words))
	for _, word := range words {
		tf[word]++
	}
	for word, freq := range tf {
		tf[word] = freq / totalWords
	}
	return tf
}

// 计算逆文档频率（IDF）
func calculateInverseDocumentFrequency(documents [][]string) map[string]float64 {
	idf := make(map[string]float64)
	totalDocuments := float64(len(documents))
	// 统计包含某个词的文档数
	wordDocumentCount := make(map[string]int)
	for _, words := range documents {
		seen := make(map[string]bool)
		for _, word := range words {
			if !seen[word] {
				wordDocumentCount[word]++
				seen[word] = true
			}
		}
	}
	// 计算逆文档频率
	for word, count := range wordDocumentCount {
		idf[word] = math.Log(totalDocuments / float64(count))
	}
	return idf
}

// 计算TF-IDF
func calculateTFIDF(words []string, documents [][]string) map[string]float64 {
	tf := calculateTermFrequency(words)
	idf := calculateInverseDocumentFrequency(documents)
	tfidf := make(map[string]float64)
	for word, freq := range tf {
		tfidf[word] = freq * idf[word]
	}
	return tfidf
}
