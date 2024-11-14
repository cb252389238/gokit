package http

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"ori/app/http/router"
	"ori/core/oriConfig"
	"ori/core/oriLog"
	"ori/core/oriSignal"
	"ori/internal/engine"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

var (
	graceful                 = flag.Bool("graceful-http", false, "listen on fd open 3 (internal use only)")
	listener    net.Listener = nil
	err         error
	reqEntityWg = &sync.WaitGroup{}
)

func Run(oriEngine *engine.OriEngine) {
	defer oriEngine.Wg.Done()
	if !flag.Parsed() {
		flag.Parse()
	}
	conf := oriConfig.GetHotConf()
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	e := router.SetupRouter(reqEntityWg)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Http.Port),
		Handler:        e,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if *graceful {
		// 子进程监听父进程传递的 socket 描述符
		oriLog.Info("平滑重启-[子进程监听文件描述]")
		// 子进程的 0, 1, 2 是预留给标准输入、标准输出、错误输出
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		oriLog.Info("正常启动-[监听信号]")
		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		oriLog.Error("listener error: %+v", err)
		return
	}
	go func() {
		if err := server.Serve(listener); err != nil {
			oriLog.Info("http服务退出,err:%+v", err)
			return
		}
	}()
	signalHandle(oriEngine, server)
}

func signalHandle(oriEngine *engine.OriEngine, server *http.Server) {
	for {
		select {
		case <-oriEngine.Context.Done():
			oriLog.Info("http服务退出")
			return
		case sig := <-oriEngine.HttpSignal:
			ctx, _ := context.WithCancel(context.Background())
			switch sig {
			case oriSignal.SIGUSR1: //优雅退出
				if err := server.Shutdown(ctx); err != nil {
					oriLog.Error(err.Error())
					break
				}
				oriLog.Info("http服务优雅退出")
				reqEntityWg.Wait() //登陆用户连接全部断开
				oriEngine.Signal <- syscall.SIGINT
			case oriSignal.SIGUSR2: //平滑重启
				err := reload() // 执行热重启函数
				if err != nil {
					oriLog.Error("http服务reload error: %v", err)
					break
				}
				oriLog.Info("http服务热重启")
				if err := server.Shutdown(ctx); err != nil {
					oriLog.Error(err.Error())
					break
				}
				reqEntityWg.Wait() //用户连接全部断开
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
