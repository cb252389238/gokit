package easy

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"
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
	int8 | int16 | int32 | int64 | int | float32 | float64
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
获取最大值
*/
type mathVal2 interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64
}

/*
返回最大值
*/
func Max[T mathVal2](nums ...T) T {
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
func Min[T mathVal2](nums ...T) T {
	if len(nums) < 2 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = T(math.Min(float64(max), float64(nums[i])))
	}
	return max
}

/*
返回随机数
*/
func Mt_rand(min, max int64) int64 {
	n, _ := crand.Int(crand.Reader, big.NewInt(max))
	r := rand.New(rand.NewSource(time.Now().UnixNano() + n.Int64()))
	return r.Int63n(max-min+1) + min
}
