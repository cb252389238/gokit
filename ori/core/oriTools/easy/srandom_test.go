package easy

import "testing"

func TestGenerateUniqueRandomNumbers(t *testing.T) {
	GenerateUniqueRandomNumbers(10, 100)
}

func BenchmarkGenerateUniqueRandomNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUniqueRandomNumbers(10, 100)
	}
}
