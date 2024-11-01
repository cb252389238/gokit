package serC

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	etcds := []string{"127.0.0.1:2379"}
	c, err := New(etcds, "/test-01", "127.0.0.1", 5, ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.WatchService("/test"))
	go func() {
		for val := range c.KeepAliveChan {
			fmt.Println("续租成功:", val)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			addService(i)
			time.Sleep(time.Second * 1)
		}
	}()
	i := 0
	for {
		/*if i == 10 {
			cancel()
		}*/
		fmt.Println("所有服务器:", c.GetServices())
		time.Sleep(time.Second * 1)
		i++
	}
}

func addService(i int) {
	etcds := []string{"127.0.0.1:2379"}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcds,
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Grant(context.Background(), 20)
	if err != nil {
		log.Fatal(err)
	}
	//注册服务并绑定租约
	_, err = cli.Put(context.Background(), "/test-0"+strconv.Itoa(i+1), "127.0.0."+strconv.Itoa(i), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
