package easy

import "testing"

func TestRandStr(t *testing.T) {
	t.Log(RandStr(10, "abcdefghijklmnopqrstuvwxyz"))
	t.Log(RandStr(10, "0123456789"))
	t.Log(RandStr(10, "0123456789abcdefghijklmnopqrstuvwxyz"))
}

func TestRandNum(t *testing.T) {
	t.Log(RandNum(100))
	t.Log(RandNum(1000000000000000000))
}

func TestN(t *testing.T) {
	t.Log(N(1, 10))
	t.Log(N(1, 1000000000000000000))
	t.Log(N(10, 20))
	t.Log(N(-10, 10))
	t.Log(N(100, 101))
}

func BenchmarkN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		N(1, 100)
	}
}
