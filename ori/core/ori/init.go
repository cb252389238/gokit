package ori

import (
	"flag"
	"fmt"
	"ori/app/http"
	"ori/app/ws"
	"ori/core/oriConfig"
	"ori/core/oriMonitor"
	"ori/core/oriSignal"
	"ori/internal/engine"
	"ori/typedef"
)

var (
	configPath string //配置文件路径
	serFlag    string
)

func Start() {
	flag.StringVar(&configPath, "f", "./config.yaml", "-f 配置文件路径")
	flag.StringVar(&serFlag, "s", "", "-s 服务1,服务2")
	if !flag.Parsed() {
		flag.Parse()
	}
	oriConfig.Load(configPath)      //载入配置文件
	go oriConfig.Listen(10)         //监听配置文件变化
	engine := engine.NewOriEngine() //初始化项目资源
	engine.Wg.Add(1)
	go oriMonitor.Monitor(engine) //监控通知
	engine.Wg.Add(1)
	go http.Run(engine) //http服务
	engine.Wg.Add(1)
	go ws.Run(engine) //websocket服务
	fmt.Println(typedef.Ico)
	fmt.Printf("服务【%s】启动完成!]\r\n", oriConfig.GetHotConf().APP)
	oriSignal.Notify(engine) //监听信号
}
