package app

import (
	"apiServer/app/router"
	"apiServer/core/coreConfig"
	"apiServer/core/coreLog"
	"apiServer/core/coreSignal"
	"apiServer/internal/engine"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

var (
	graceful              = flag.Bool("graceful-http", false, "listen on fd open 3 (internal use only)")
	listener net.Listener = nil
	err      error
	wg       = &sync.WaitGroup{}
)

func Run(oriEngine *engine.OriEngine) {
	defer oriEngine.Wg.Done()
	if !flag.Parsed() {
		flag.Parse()
	}
	conf := coreConfig.GetHotConf()
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	e := router.SetupRouter(oriEngine, wg)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Http.Port),
		Handler:        e,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if *graceful {
		// 子进程监听父进程传递的 socket 描述符
		coreLog.LogInfo("平滑重启-[子进程监听文件描述]")
		// 子进程的 0, 1, 2 是预留给标准输入、标准输出、错误输出
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		coreLog.LogInfo("正常启动-[监听信号]")
		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		coreLog.LogError("listener error: %+v", err)
		return
	}
	go func() {
		if err := server.Serve(listener); err != nil {
			coreLog.LogInfo("http服务退出,err:%+v", err)
			return
		}
	}()
	signalHandle(oriEngine, server)
}

func signalHandle(oriEngine *engine.OriEngine, server *http.Server) {
	for {
		select {
		case <-oriEngine.Context.Done():
			coreLog.LogInfo("http服务退出")
			return
		case sig := <-oriEngine.HttpSignal:
			ctx, _ := context.WithCancel(context.Background())
			switch sig {
			case coreSignal.SIGUSR1: //优雅退出
				if err := server.Shutdown(ctx); err != nil {
					coreLog.LogError(err.Error())
					break
				}
				coreLog.LogInfo("http服务优雅退出")
				wg.Wait() //登陆用户连接全部断开
				oriEngine.Signal <- syscall.SIGINT
			case coreSignal.SIGUSR2: //平滑重启
				err := reload() // 执行热重启函数
				if err != nil {
					coreLog.LogError("http服务reload error: %v", err)
					break
				}
				coreLog.LogInfo("http服务热重启")
				if err := server.Shutdown(ctx); err != nil {
					coreLog.LogError(err.Error())
					break
				}
				wg.Wait() //用户连接全部断开
				oriEngine.Signal <- syscall.SIGINT
			}
		}
	}
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}
	// 获取 socket 描述符
	f, err := tl.File()
	if err != nil {
		return err
	}
	// 设置传递给子进程的参数（包含 socket 描述符）
	args := []string{"-graceful-http"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} // 文件描述符
	// 新建并执行子进程
	return cmd.Start()
}
