package centrifuge

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func NewCentrifuge() {
	err := New([]string{"pc", "web", "app"}, time.Second*10)
	if err != nil {
		panic(err)
	}
}

func TestProcess(t *testing.T) {
	NewCentrifuge()
	Instance().CacheDeleteCallback(func(k string, v any) {
		fmt.Println("触发回调缓存回调-", k, v)
	})
	time.Sleep(time.Second)
	sub1, err := Instance().Add("1001", "pc", nil)
	if err != nil {
		log.Fatal("1001加入失败")
	}
	sub2, err := Instance().Add("1002", "web", nil)
	if err != nil {
		log.Fatal("1002加入失败")
	}
	sub3, err := Instance().Add("1003", "app", nil)
	if err != nil {
		log.Fatal("1003加入失败")
	}
	fmt.Printf("在线人数:%d,在线连接数:%d\r\n", Instance().GetUserNum(), Instance().GetConnectNum())
	//心跳模拟
	go func() {
		for {
			time.Sleep(time.Second * 5)
			Instance().Heartbeat("1001", "pc")
			Instance().Heartbeat("1002", "web")
			Instance().Heartbeat("1003", "app")
		}
	}()
	//模拟获取消息
	go func() {
		for {
			select {
			case msg := <-sub1.C:
				fmt.Println("1001 pc-收到消息：", msg)
			case msg := <-sub2.C:
				fmt.Println("1002 web-收到消息：", msg)
			case msg := <-sub3.C:
				fmt.Println("1003 app-收到消息：", msg)
			}
		}
	}()
	fmt.Println("1001,1002,1003订阅频道1")
	err = Instance().Subscribe("1001", "pc", "1")
	if err != nil {
		log.Fatal("1001 pc", err)
	}
	err = Instance().Subscribe("1002", "web", "1")
	if err != nil {
		log.Fatal("1002 web", err)
	}
	err = Instance().Subscribe("1003", "app", "1")
	if err != nil {
		log.Fatal("1003 app", err)
	}
	fmt.Println("向频道1推送消息 1001,1002,1003收到消息")
	Instance().PushToTopic("1", "hello world")
	time.Sleep(time.Second * 3)
	fmt.Println("向频道1的pc用户推送消息 1001收到消息")
	Instance().PushToTopicByPlatform("1", "hello world", "pc")
	time.Sleep(time.Second * 3)
	fmt.Println("1001 pc删除")
	err = Instance().Del("1001", "pc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("向频道1的pc用户推送消息 1001收不到消息")
	Instance().PushToTopicByPlatform("1", "hello world", "pc")
	time.Sleep(time.Second * 3)
	fmt.Println("向频道1推送消息 1002,1003收到消息")
	Instance().PushToTopicByPlatform("1", "hello world", "web", "app")
	time.Sleep(time.Second * 3)
	fmt.Printf("在线人数:%d,在线连接数:%d\r\n", Instance().GetUserNum(), Instance().GetConnectNum())
	Instance().Add("1004", "pc", nil)
	Instance().Add("1004", "web", nil)
	Instance().Add("1004", "app", nil)
	fmt.Printf("在线人数:%d,在线连接数:%d\r\n", Instance().GetUserNum(), Instance().GetConnectNum())
	select {}
}
