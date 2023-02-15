package ori

import (
	"fmt"
	"ori/app/http"
	"ori/internal"
	"ori/internal/core/config"
	"ori/internal/core/monitor"
	"ori/internal/core/oriEngine"
	"ori/internal/core/oriSignal"
	"ori/typedef"
	"time"
)

func Start() {
	config.Load()                      //载入配置文件
	go config.Listen(10)               //监听配置文件变化
	engine := oriEngine.NewOriEngine() //初始化项目资源
	engine.Wg.Add(1)
	go monitor.Monitor(engine) //监控通知
	//每天十点报告状态
	go func() {
		for {
			now := time.Now()
			next := now.Add(24 * time.Hour)
			next = time.Date(next.Year(), next.Month(), next.Day(), 10, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
	engine.Wg.Add(1)
	go internal.Run(engine) //核心逻辑
	engine.Wg.Add(1)
	go http.Run(engine)
	fmt.Println(typedef.Ico)
	oriSignal.Notify(engine) //监听信号
}
