package bloomfilter

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	val := 0
	b := NewBloomFilter(10000000, 0.0001)
	for i := 10000; i < 10010000; i++ {
		b.Add([]byte(strconv.Itoa(i)))
	}
	for i := 10000; i < 10010000; i++ {
		if b.Contains([]byte(strconv.Itoa(i))) {
			val++
		}
	}
	fmt.Println(val)
}

func BenchmarkNewBloomFilter(b *testing.B) {
	bloom := NewBloomFilter(10000000, 0.0001)
	//for i := 10000; i < 10010000; i++ {
	//	bloom.Add([]byte(strconv.Itoa(i)))
	//}
	for i := 0; i < b.N; i++ {
		bloom.Contains([]byte(strconv.Itoa(i)))
	}
}
