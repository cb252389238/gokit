package easy

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// 高性能随机数生成器（分片版）
var randPool = struct {
	sources []*rand.Rand
	count   int
}{
	sources: make([]*rand.Rand, 32), // 根据CPU核心数调整
	count:   0,
}

func init() {
	seed := time.Now().UnixNano()
	for i := range randPool.sources {
		randPool.sources[i] = rand.New(rand.NewSource(seed + int64(i)))
	}
}

// GenerateUniqueRandomNumbersV2 高性能不重复随机数生成器
// 长度超过 1e6 时建议使用 crypto/rand 方案
func GenerateUniqueRandomNumbers(length int, maxRange ...int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("length must be greater than 0")
	}

	// 自动计算安全范围
	effectiveRange := math.MaxInt
	if len(maxRange) > 0 && maxRange[0] > 0 {
		effectiveRange = maxRange[0]
	}

	// 冲突概率检查（生日问题）
	if float64(length) > 0.5*math.Sqrt(float64(effectiveRange)) {
		return nil, errors.New("high collision probability, consider increasing range")
	}

	// 选择最优算法
	if effectiveRange >= length*1000 {
		return lcgAlgorithm(length, effectiveRange)
	}
	return mapAlgorithm(length, effectiveRange)
}

// 线性同余算法（适合大范围）
func lcgAlgorithm(length, max int) ([]int, error) {
	// 生成互质参数
	a, c := uint64(6364136223846793005), uint64(1442695040888963407)
	seed := uint64(time.Now().UnixNano())

	result := make([]int, length)
	exists := make(map[int]struct{}, length)

	for i := 0; i < length; {
		seed = seed*a + c
		num := int(seed % uint64(max))

		if _, ok := exists[num]; !ok {
			exists[num] = struct{}{}
			result[i] = num
			i++
		}
	}
	return result, nil
}

// 哈希查重算法（适合小范围）
func mapAlgorithm(length, max int) ([]int, error) {
	// 使用分片随机源
	index := randPool.count % len(randPool.sources)
	r := randPool.sources[index]
	randPool.count++

	result := make([]int, 0, length)
	exists := make(map[int]struct{}, length)

	for len(result) < length {
		num := r.Intn(max)
		if _, ok := exists[num]; !ok {
			exists[num] = struct{}{}
			result = append(result, num)
		}
	}
	return result, nil
}
