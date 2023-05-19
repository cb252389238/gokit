package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

var (
	m  = sync.Map{}
	wg sync.WaitGroup
)

func TestID_Node(t *testing.T) {
	for nnum := 1; nnum < 11; nnum++ {
		wg.Add(1)
		go func(nnum int64) {
			defer wg.Done()
			node, err := NewNode(nnum)
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := 0; i <= 10; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					for j := 0; j <= 100000; j++ {
						i2 := node.Generate().Int64()
						if value, ok := m.Load(i2); ok {
							m.Store(i2, value.(int64)+1)
						} else {
							m.Store(i2, int64(1))
						}
					}
				}()
			}
		}(int64(nnum))
	}

	wg.Wait()
	m.Range(func(key, value any) bool {
		if value.(int64) > 1 {
			fmt.Println(key)
		}
		return true
	})
}
