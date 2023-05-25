package pubsub

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestPubsub(t *testing.T) {
	//初始化一个发布者对象
	pub := New()
	user1Sub, _ := pub.Subscribe("topic1", "1001")
	user1Sub, _ = pub.Subscribe("topic2", "1001")
	user2Sub, _ := pub.Subscribe("topic1", "1002")

	go func() {
		i := 0
		for {
			pub.Publish("topic1", "topic1测试消息"+strconv.Itoa(i))
			i++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		i := 0
		for {
			pub.Publish("topic2", "topic2测试消息"+strconv.Itoa(i))
			i++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for message := range user1Sub.C {
			fmt.Println("user1-" + message.(string))
		}
	}()
	go func() {
		for message := range user2Sub.C {
			fmt.Println("user2-" + message.(string))
		}
	}()
	pub.PublishToUser("1001", "单独推送给用户的消息")
	pub.PublishAll("推送给所有用户的消息")
	for {
		select {
		case <-time.After(time.Second * 10):
			pub.Exit(user2Sub, "topic1")
			pub.Exit(user1Sub, "topic1")
			time.Sleep(time.Second * 5)
			return
		}
	}
}
