package pubsub

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestPubsub(t *testing.T) {
	//初始化一个发布者对象
	pub, err := NewPublisher(3)
	if err != nil {
		panic(err)
	}
	user1 := pub.SubscriberTopic("1001")
	user2 := pub.SubscriberTopic("1001")
	user3 := pub.SubscriberTopic("1002")
	user4 := pub.SubscriberTopic("1003")

	go func() {
		for i := 0; i < 10; i++ {
			pub.Publish("1001", "1001测试消息"+strconv.Itoa(i))
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			pub.Publish("1002", "1002测试消息"+strconv.Itoa(i))
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for i := 0; i < 5; i++ {
			pub.Publish("1003", "1003测试消息"+strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for message := range user1.C {
			fmt.Println("user1-" + message.(string))
		}
		fmt.Println("user1-退出订阅")
	}()
	go func() {
		for message := range user2.C {
			fmt.Println("user2-" + message.(string))
			time.Sleep(time.Second)
		}
		fmt.Println("user2-退出订阅")
	}()
	go func() {
		for message := range user3.C {
			fmt.Println("user3-" + message.(string))
		}
		fmt.Println("user3-退出订阅")
	}()
	go func() {
		for message := range user4.C {
			fmt.Println("user4-" + message.(string))
		}
		fmt.Println("user4-退出订阅")
	}()
	for {
		select {
		case <-time.After(time.Second * 5):
			pub.Exict(user1)
			time.Sleep(time.Second * 10)
			return
		}
	}
}
