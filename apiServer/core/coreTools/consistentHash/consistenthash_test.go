package consistentHash

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	sets = []string{"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4", "127.0.0.5"}
	res  = map[string]int{}
)

func TestHashing(t *testing.T) {
	hash := New(1000)
	hash.Set(sets)
	for uid := 1; uid <= 10000000; uid++ {
		key, _ := hash.Get(strconv.Itoa(uid))
		res[key] += 1
	}
	fmt.Println(res)
}

func BenchmarkHashing(b *testing.B) {
	hash := New(1000)
	hash.Set(sets)
	for i := 0; i < b.N; i++ {
		hash.Get(strconv.Itoa(i))
	}
}
