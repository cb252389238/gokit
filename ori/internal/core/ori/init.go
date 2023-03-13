package ori

import (
	"flag"
	"fmt"
	"ori/internal/core/config"
	"ori/internal/core/monitor"
	"ori/internal/core/oriEngine"
	"ori/internal/core/oriSignal"
	"ori/internal/service"
	"ori/typedef"
	"time"
)

var (
	configPath string //配置文件路径
)

func Start() {
	flag.StringVar(&configPath, "f", "./config.yaml", "-f 配置文件路径")
	flag.Parse()
	config.Load(configPath) //载入配置文件
	//return
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
	service.Run(engine) //服务启动
	fmt.Println(typedef.Ico)
	oriSignal.Notify(engine) //监听信号
}
