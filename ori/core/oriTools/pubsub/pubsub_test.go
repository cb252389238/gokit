package pubsub

import (
	"fmt"
	"testing"
	"time"
)

func TestPubsub(t *testing.T) {
	//初始化一个发布者对象
	platformArr := []string{"pc", "h5", "app"}
	pub, err := New(platformArr)
	if err != nil {
		t.Fatal(err)
	}
	user1_pc, _ := pub.Subscribe("topic1", "1001", "pc")
	user1_app, _ := pub.Subscribe("topic1", "1001", "app")
	user2_app, _ := pub.Subscribe("topic1", "1002", "app")
	user3_app, _ := pub.Subscribe("topic2", "1003", "app")
	user4_app, _ := pub.Subscribe("topic2", "1004", "app")

	go func() {
		for message := range user1_pc.C {
			fmt.Println("1001-pc:" + message.(string))
		}
	}()
	go func() {
		for message := range user1_app.C {
			fmt.Println("1001-app:" + message.(string))
		}
	}()
	go func() {
		for message := range user2_app.C {
			fmt.Println("1002-app:" + message.(string))
		}
	}()
	go func() {
		for message := range user3_app.C {
			fmt.Println("1003-app:" + message.(string))
		}
	}()
	go func() {
		for message := range user4_app.C {
			fmt.Println("1004-app:" + message.(string))
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("向topic1推送消息，1001-pc,1001-app，1002-app收到消息")
	pub.Publish("topic1", "推送给订阅topic1的用户的消息")
	time.Sleep(time.Second)
	fmt.Println("向topic2推送消息，1003-app，1004-app收到消息")
	pub.Publish("topic2", "推送给订阅topic2的用户的消息")
	time.Sleep(time.Second)
	fmt.Println("向所有用户推送消息 1001-pc,1001-app，1002-app，1003-app，1004-app收到消息")
	pub.PublishAll("推送给所有用户的消息")
	time.Sleep(time.Second)
	fmt.Println("向1001推送消息 1001-pc,1001-app收到")
	pub.PublishToUser("1001", "推送给1001用户的消息")
	time.Sleep(time.Second)
	fmt.Println("向1001app推送消息 1001-app收到")
	pub.PublishToUser("1001", "推送给1001-app用户的消息", "app")
	time.Sleep(time.Second)
	fmt.Println("向1001pc推送消息 1001-pc收到")
	pub.PublishToUser("1001", "推送给1001-pc用户的消息", "pc")
	time.Sleep(time.Second)
	fmt.Println("1001-pc退出订阅")
	pub.Exit("topic1", "1001", "pc")
	time.Sleep(time.Second)
	fmt.Println("向topic1推送消息，1001-app，1002-app收到消息")
	pub.Publish("topic1", "推送给订阅topic1的用户的消息")
}
