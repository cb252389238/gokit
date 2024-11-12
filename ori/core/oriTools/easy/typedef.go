package easy

type PHP_ARRAY interface {
	[]uint | []uint8 | []uint16 | []uint32 |
		[]uint64 | []int | []int8 | []int16 |
		[]int32 | []int64 |
		[]uintptr | []float32 | []float64 |
		[]complex64 | []complex128 |
		[]string | []any |
		map[string]any |
		map[int]any
}

type phpArray interface {
	map[int]int | map[int]string | map[string]int | map[string]string | []int | []uint | []string
}
