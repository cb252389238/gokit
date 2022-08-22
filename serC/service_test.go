package serC

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	etcds := []string{"127.0.0.1:2379"}
	c, err := New(etcds, "/test-01", "127.0.0.1", 5, ctx)
	fmt.Println(c.WatchService("/test"))
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for val := range c.KeepAliveChan {
			fmt.Println("续租成功:", val)
		}
	}()
	go func() {
		time.Sleep(time.Second * 3)
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   etcds,
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			log.Fatal(err)
		}
		resp, err := cli.Grant(context.Background(), 5)
		if err != nil {
			log.Fatal(err)
		}
		//注册服务并绑定租约
		_, err = cli.Put(context.Background(), "/test-02", "127.0.0.2", clientv3.WithLease(resp.ID))
		if err != nil {
			log.Fatal(err)
		}
	}()
	i := 0
	for {
		if i == 10 {
			cancel()
		}
		fmt.Println("所有服务器:", c.GetServices())
		time.Sleep(time.Second * 1)
		i++
	}
}
