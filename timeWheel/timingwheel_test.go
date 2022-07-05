package timeWheel

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTimingWheel(t *testing.T) {
	interval := time.Second * 1 //时间单位
	numSlots := 60              //时间槽数
	//到时间得执行回调函数
	execute := func(v interface{}) {
		fmt.Println(v)
	}
	wheel, err := NewTimingWheel(interval, numSlots, execute)
	if err != nil {
		panic(err)
	}
	id := ""
	for i := 1; i <= 5000000; i++ {
		t := i % 10
		if t == 0 {
			t = 1
		}
		if t == 9 {
			id = wheel.SetTimer("val"+strconv.Itoa(i), time.Second*time.Duration(i%10))
			wheel.RemoveTimer(id)
		} else {
			wheel.SetTimer("val"+strconv.Itoa(i), time.Second*time.Duration(i%10))
		}
	}
	fmt.Println("写入成功")
	for {

	}
}

func BenchmarkNewTimingWheel(b *testing.B) {
	interval := time.Second * 1
	numSlots := 60
	execute := func(v interface{}) {
		//fmt.Println(v)
	}
	wheel, err := NewTimingWheel(interval, numSlots, execute)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		wheel.SetTimer("val"+strconv.Itoa(i), time.Second*time.Duration(i%10000))
	}
}
