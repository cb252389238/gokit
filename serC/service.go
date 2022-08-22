package serC

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 服务注册发现
type SerC struct {
	cli           *clientv3.Client
	endpoints     []string                                //etcd节点
	serverList    map[string]string                       //服务列表
	leaseID       clientv3.LeaseID                        //租约ID
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse //租约keepalieve相应chan
	lock          sync.RWMutex
	lease         int64 //租约生命时间
	ctx           context.Context
}

// etcd服务节点
func New(endpoints []string, key, val string, lease int64, ctx context.Context) (*SerC, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	serc := &SerC{cli: cli, endpoints: endpoints, lease: lease, ctx: ctx}
	serc.serverList = make(map[string]string)
	if err := serc.putKeyWithLease(key, val); err != nil {
		return nil, err
	}
	return serc, nil
}

// 设置租约
func (s *SerC) putKeyWithLease(key, val string) error {
	//设置租约时间
	resp, err := s.cli.Grant(s.ctx, s.lease)
	if err != nil {
		return err
	}
	//注册服务并绑定租约
	_, err = s.cli.Put(s.ctx, key, val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	leaseRespChan, err := s.cli.KeepAlive(s.ctx, resp.ID)
	if err != nil {
		return err
	}
	s.KeepAliveChan = leaseRespChan
	s.leaseID = resp.ID
	return nil
}

// 关闭租约
func (s *SerC) Close() error {
	//撤销租约
	if _, err := s.cli.Revoke(s.ctx, s.leaseID); err != nil {
		return err
	}
	return nil
}

// WatchService 初始化服务列表和监视
func (s *SerC) WatchService(prefix string) error {
	//根据前缀获取现有的key
	resp, err := s.cli.Get(s.ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, ev := range resp.Kvs {
		s.setServiceList(string(ev.Key), string(ev.Value))
	}
	//监视前缀，修改变更的server
	go s.watcher(prefix)
	return nil
}

// watcher 监听前缀
func (s *SerC) watcher(prefix string) {
	rch := s.cli.Watch(s.ctx, prefix, clientv3.WithPrefix())
	for {
		select {
		case <-s.ctx.Done():
			fmt.Println("退出服务发现")
			return
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				switch ev.Type {
				case mvccpb.PUT: //修改或者新增
					fmt.Println("新增/更新服务器")
					s.setServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
				case mvccpb.DELETE: //删除
					fmt.Println("删除服务器")
					s.delServiceList(string(ev.Kv.Key))
				}
			}
		}
	}
}

// SetServiceList 新增服务地址
func (s *SerC) setServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = string(val)
}

// DelServiceList 删除服务地址
func (s *SerC) delServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
}

// GetServices 获取服务地址
func (s *SerC) GetServices() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)
	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}
