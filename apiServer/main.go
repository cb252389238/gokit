//go:generate 	swag init --parseDependency --parseDepth=6 --exclude ./core,./internal/,./util/  -o ./docs  --instanceName yfapi
package main

import (
	"apiServer/app"
	"apiServer/core/coreConfig"
	"apiServer/core/coreSignal"
	eng "apiServer/internal/engine"
	"apiServer/typedef"
	"apiServer/util/easy"
	"flag"
	"fmt"
	"log"
)

// @title		api文档
// @version		1.0
// @description	api文档
// @license.name	Apache 2.0
// @contact.name	api文档
// @host			localhost:8001
// @BasePath /api

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	var configPath string //配置文件路径
	flag.StringVar(&configPath, "f", "./config/config.yaml", "-f 配置文件路径")
	if !flag.Parsed() {
		flag.Parse()
	}
	coreConfig.Load(configPath)  //载入配置文件
	go coreConfig.Listen(10)     //监听配置文件变化
	engine := eng.NewOriEngine() //初始化项目资源
	engine.Wg.Add(1)
	go app.Run(engine) //http服务
	fmt.Println(easy.Green(typedef.Ico))
	fmt.Printf(easy.Green("服务【%s:%d】启动完成!]\r\n"), coreConfig.GetHotConf().APP, coreConfig.GetHotConf().Http.Port)
	coreSignal.Notify(engine) //监听信号
}
