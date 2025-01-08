package easy

import (
	"crypto/rand"
	"math"
	"math/big"
)

// 根据坐标计算角度
func Angle(x, y float64) float64 {
	r := (math.Atan2(y, x) * 180) / math.Pi
	if r < 0 {
		return 360 + r
	}
	return r
}

type mathVal interface {
	int8 | int16 | int32 | int64 | int | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

/*
绝对值
*/
func Abs[T mathVal](number T) T {
	return T(math.Abs(float64(number)))
}

/*
向上取整
*/
func Ceil[T float32 | float64](number T) T {
	return T(math.Ceil(float64(number)))
}

/*
向下取整
*/
func Floor[T float32 | float64](number T) T {
	return T(math.Floor(float64(number)))
}

/*
返回最大值
*/
func Max[T mathVal](nums ...T) T {
	if len(nums) < 2 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = T(math.Max(float64(max), float64(nums[i])))
	}
	return max
}

/*
返回最小值
*/
func Min[T mathVal](nums ...T) T {
	if len(nums) < 2 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = T(math.Min(float64(max), float64(nums[i])))
	}
	return max
}

// isPrime 判断一个数字是否是质数
func IsPrime(n *big.Int) bool {
	return n.ProbablyPrime(0)
}

// generateRandomPrime 生成指定位数的随机质数
func GenerateRandomPrime(bitLength int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bitLength)
	if err != nil {
		return nil, err
	}
	return prime, nil
}
