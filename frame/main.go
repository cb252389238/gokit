package main

import (
	"fmt"
	"frame/internal"
	"frame/internal/svc"
	"frame/resource"
	"frame/types"
	"os"
	"os/signal"
	"syscall"
)

// 监听信号
func signalHandler(ctx *svc.ServiceContext, s os.Signal) {
	switch s {
	case syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT: //强制退出，不会优雅关闭，立即退出程序 kill pid,kill -9 pid
		ctx.Cancle()
		ctx.Wg.Wait() //阻塞等待同步协程完成已载入任务
		resource.LogInfo("立刻退出程序。信号:[%s]", s.String())
		os.Exit(0)
	case types.SIGUSR1: //优雅关闭，处理完所有连接退出 kill -USR1 [PID]
		resource.LogInfo("优雅退出程序。信号:[%s]", s.String())
		ctx.WsC <- types.SIGUSR1
	case types.SIGUSR2: //平滑重启信号 kill -USR2 [PID]
		resource.LogInfo("平滑重启程序。信号:[%s]", s.String())
		ctx.WsC <- types.SIGUSR2
	default:
		fmt.Printf("未定义信号忽略！信号:[%s]\r\n", s.String())
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	resource.LoadConfig() //载入配置文件
	go resource.ListenConfig()
	ctx := svc.NewServiceContext() //初始化项目资源
	ctx.Wg.Add(1)
	go internal.Run(ctx) //核心逻辑
	fmt.Println(types.Ico)
	signal.Notify(ctx.C, types.Signals...)
	for s := range ctx.C {
		signalHandler(ctx, s) //退出服务
	}
}
