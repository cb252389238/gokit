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
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	var configPath string //配置文件路径
	flag.StringVar(&configPath, "f", "./config.yaml", "-f 配置文件路径")
	if !flag.Parsed() {
		flag.Parse()
	}
	oriConfig.Load(configPath)         //载入配置文件
	go oriConfig.Listen(10)            //监听配置文件变化
	engine := oriEngine.NewOriEngine() //初始化项目资源
	engine.Wg.Add(1)
	go oriMonitor.Monitor(engine) //监控通知
	engine.Wg.Add(1)
	go http.Run(engine) //http服务
	engine.Wg.Add(1)
	go ws.Run(engine) //websocket服务
	engine.Wg.Add(1)
	service.Run(engine) //自定义服务
	fmt.Println(typedef.Ico)
	fmt.Printf("服务【%s】启动完成!]\r\n", oriConfig.GetHotConf().APP)
	oriSignal.Notify(engine) //监听信号
}
