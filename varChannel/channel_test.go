package varChannel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	c := New()
	go func() {
		for i := 0; ; i++ {
			if err := c.Write(i); err != nil {
				fmt.Println(err)
			} else {
				time.Sleep(time.Second)
			}
		}
	}()
	go func() {
		time.Sleep(time.Second * 10)
		c.Close()
	}()
	for v := range c.Read() {
		fmt.Println(v)
	}
	fmt.Println("end")
}
