package ori

import (
	"flag"
	"fmt"
	"ori/app/http"
	"ori/app/ws"
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
	services := config.GetHotConf().Services
	if services.CONFIG_HOT_UPDATE_SERVER {
		go config.Listen(10) //监听配置文件变化
	}
	engine := oriEngine.NewOriEngine() //初始化项目资源
	if services.MONITOR_SERVER {
		engine.Wg.Add(1)
		go monitor.Monitor(engine) //监控通知
	}
	if services.STATUS_REPORT_SERVER {
		//每天报告状态
		go func() {
			for {
				now := time.Now()
				next := now.Add(24 * time.Hour)
				next = time.Date(next.Year(), next.Month(), next.Day(), config.GetHotConf().StatusReportHour, 0, 0, 0, next.Location())
				t := time.NewTimer(next.Sub(now))
				<-t.C
			}
		}()
	}
	if services.HTTP_SERVER {
		engine.Wg.Add(1)
		go http.Run(engine)
	}
	if services.WEBSOCKET_SERVER {
		engine.Wg.Add(1)
		go ws.Run(engine)
	}

	service.Run(engine) //自定义服务
	fmt.Println(typedef.Ico)
	fmt.Printf("服务【%s】启动完成!]", config.GetHotConf().APP)
	oriSignal.Notify(engine) //监听信号
}
