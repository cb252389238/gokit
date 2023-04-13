package bloomfilter

import (
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	bitSet  []bool
	hashNum int
}

// n 预期的数据量， p期望判断失误率
func NewBloomFilter(n uint, p float64) *BloomFilter {
	m := uint(math.Ceil(float64(n) * math.Log(p) / math.Log(1.0/(math.Pow(2.0, math.Log(2.0))))))
	k := uint(math.Ceil(math.Log(2.0) * float64(m) / float64(n)))
	return &BloomFilter{
		bitSet:  make([]bool, m),
		hashNum: int(k),
	}
}

// 生成hash值
func (bf *BloomFilter) hashFuncs(key []byte) []uint {
	h := fnv.New64()
	h.Write(key)
	hashVal := h.Sum64()
	var hashVals []uint
	for i := 0; i < bf.hashNum; i++ {
		hashVals = append(hashVals, uint(hashVal))
		hashVal += hashVal
	}
	return hashVals
}

// 添加到布隆过滤器内
func (bf *BloomFilter) Add(key []byte) {
	for _, i := range bf.hashFuncs(key) {
		bf.bitSet[i%uint(len(bf.bitSet))] = true
	}
}

// 判断数据是否存在 true存在 false不存在
func (bf *BloomFilter) Contains(key []byte) bool {
	for _, i := range bf.hashFuncs(key) {
		if !bf.bitSet[i%uint(len(bf.bitSet))] {
			return false
		}
	}
	return true
}
