package minhash

import (
	"hash/fnv"
	"math"
)

const (
	numHashFunctions = 100 // MinHash函数数量
)

// 计算单个文档的MinHash签名
func ComputeMinHashSignature(document []string) []uint32 {
	signature := make([]uint32, numHashFunctions)

	// 初始化所有哈希值为正无穷大
	for i := range signature {
		signature[i] = math.MaxUint32
	}

	// 使用哈希函数计算MinHash签名
	for _, word := range document {
		hash := hashString(word)
		for i := 0; i < numHashFunctions; i++ {
			hashValue := computeHashValue(hash, i)
			if hashValue < signature[i] {
				signature[i] = hashValue
			}
		}
	}

	return signature
}

// 计算两个MinHash签名的相似度
func ComputeSimilarity(signature1, signature2 []uint32) float64 {
	if len(signature1) != len(signature2) {
		panic("MinHash signatures must have the same length")
	}
	matches := 0
	for i := 0; i < len(signature1); i++ {
		if signature1[i] == signature2[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(signature1))
}

// 哈希字符串
func hashString(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

// 计算哈希值
func computeHashValue(hash uint32, i int) uint32 {
	a := uint32(i*2 + 1)
	b := uint32(i*2 + 2)
	return (hash*a + b) % math.MaxUint32
}
