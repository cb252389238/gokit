package imMerge

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestMerge(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	engine := New(Config{
		MergeThreshold: 10,
		MergeTime:      200,
		TimeOut:        1000,
		ctx:            ctx,
	})
	for j := 0; j < 100; j++ {
		go func(j int) {
			for i := 0; i <= 10000; i++ {
				engine.Message(strconv.Itoa(j), i)
			}
		}(j)
	}

	for v := range engine.Result {
		fmt.Printf("groupid:%s,data:%+v\r\n", v.groupId, v.data)
	}
}
