package internal

import (
	"frame/internal/monitor"
	"frame/internal/svc"
	"frame/resource"
	"time"
)

func Run(ctx *svc.ServiceContext) {
	defer ctx.Wg.Done()
	go resource.ListenConfig() //监听配置文件变化
	//ctx.Wg.Add(1)
	//go ws.Run(ctx) //websocket协程
	ctx.Wg.Add(1)
	go monitor.MonitorService(ctx) //监控通知
	go func() {
		for {
			now := time.Now()
			next := now.Add(24 * time.Hour)
			next = time.Date(next.Year(), next.Month(), next.Day(), 10, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
