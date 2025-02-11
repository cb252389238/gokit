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

// 地球半径，单位为米
const earthRadius = 6371000

// 计算两个经纬度坐标之间的距离
func HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// 将角度转换为弧度
	radLat1 := lat1 * math.Pi / 180
	radLon1 := lon1 * math.Pi / 180
	radLat2 := lat2 * math.Pi / 180
	radLon2 := lon2 * math.Pi / 180

	// 计算经纬度差值
	deltaLat := radLat2 - radLat1
	deltaLon := radLon2 - radLon1

	// 使用Haversine公式计算距离
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(radLat1)*math.Cos(radLat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// 计算距离
	distance := earthRadius * c

	return distance
}
