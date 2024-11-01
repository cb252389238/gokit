package safemap

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestSafeMap(T *testing.T) {
	m := NewSafeMap()
	wg.Add(1)
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Set(i, i)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Get(i)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 0; i < 10000000; i++ {
			m.Get(i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end")
}

func BenchmarkSafeMap_Set(b *testing.B) {
	m := NewSafeMap()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
}

func BenchmarkSafeMap_Get(b *testing.B) {
	m := NewSafeMap()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
}
