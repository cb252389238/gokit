package easy

import "math"

// 根据坐标计算角度
func Angle(x, y float64) float64 {
	r := (math.Atan2(y, x) * 180) / math.Pi
	if r < 0 {
		return 360 + r
	}
	return r
}
