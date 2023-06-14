package main

import (
	"flag"
	"fmt"
	"log"
	"ori/app/http"
	"ori/app/ws"
	"ori/core/oriConfig"
	"ori/core/oriEngine"
	"ori/core/oriMonitor"
	"ori/core/oriSignal"
	"ori/internal/service"
	"ori/typedef"
	"time"
)

func start() {
	var configPath string //配置文件路径
	var serFlag string
	flag.StringVar(&configPath, "f", "./config.yaml", "-f 配置文件路径")
	flag.StringVar(&serFlag, "s", "", "-s 服务1,服务2")
	if !flag.Parsed() {
		flag.Parse()
	}
	oriConfig.Load(configPath)         //载入配置文件
	go oriConfig.Listen(10)            //监听配置文件变化
	engine := oriEngine.NewOriEngine() //初始化项目资源
	engine.Wg.Add(1)
	go oriMonitor.Monitor(engine) //监控通知
	//每天报告状态
	go func() {
		for {
			now := time.Now()
			next := now.Add(24 * time.Hour)
			next = time.Date(next.Year(), next.Month(), next.Day(), oriConfig.GetHotConf().StatusReportHour, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
	engine.Wg.Add(1)
	go http.Run(engine) //http服务
	engine.Wg.Add(1)
	go ws.Run(engine) //websocket服务
	engine.Wg.Add(1)
	service.Run(engine, serFlag) //自定义服务
	fmt.Println(typedef.Ico)
	fmt.Printf("服务【%s】启动完成!]\r\n", oriConfig.GetHotConf().APP)
	oriSignal.Notify(engine) //监听信号
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	start() //启动项目
}
