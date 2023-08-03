package cosinesim

import "math"

// 计算词频（TF）
func CalculateTermFrequency(words []string) map[string]float64 {
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

// 计算余弦相似度
func ComputeCosineSimilarity(tf1, tf2 map[string]float64) float64 {
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	// 计算点积和向量的长度
	for word, freq1 := range tf1 {
		freq2, exists := tf2[word]
		dotProduct += freq1 * freq2
		magnitude1 += freq1 * freq1
		if exists {
			magnitude2 += freq2 * freq2
		}
	}

	// 处理不存在于tf1但存在于tf2的词
	for word, freq2 := range tf2 {
		if _, exists := tf1[word]; !exists {
			magnitude2 += freq2 * freq2
		}
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0.0
	}

	return dotProduct / (magnitude1 * magnitude2)
}
