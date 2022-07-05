package safelist

import (
	"fmt"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestList(t *testing.T) {
	list := New()
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			list.PushBack(i)
			list.PushFront(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end")
}
