package oriSignal

import (
	"ori/internal/core/log"
	"ori/internal/core/oriEngine"
	"os"
	"os/signal"
	"syscall"
)

func listen(engine *oriEngine.OriEngine, s os.Signal) {
	switch s {
	case syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT: //强制退出，立即退出程序 kill pid,kill -9 pid
		engine.Cancel()
		engine.Wg.Wait() //阻塞等待同步协程完成
		log.LogInfo("立刻退出程序。信号:[%s]", s.String())
		os.Exit(0)
	case SIGUSR1: //优雅关闭，处理完所有连接退出 kill -USR1 [PID]
		log.LogInfo("优雅退出程序。信号:[%s]", s.String())
		//engine.WsSignal <- SIGUSR1
		engine.HttpSignal <- SIGUSR1
	case SIGUSR2: //平滑重启信号 kill -USR2 [PID]
		log.LogInfo("平滑重启程序。信号:[%s]", s.String())
		//engine.WsSignal <- SIGUSR2
		engine.HttpSignal <- SIGUSR2
	default:
		log.LogInfo("未定义信号忽略！信号:[%s]", s.String())
	}
}

func Notify(engine *oriEngine.OriEngine) {
	signal.Notify(engine.Signal, Signals...)
	for s := range engine.Signal {
		listen(engine, s) //退出服务
	}
}