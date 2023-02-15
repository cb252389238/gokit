package http

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"ori/internal/core/config"
	"ori/internal/core/log"
	"ori/internal/core/oriEngine"
	"ori/internal/core/oriSignal"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

var (
	graceful              = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
	listener net.Listener = nil
	err      error
	wg       = &sync.WaitGroup{}
)

func handle(oriEngine *oriEngine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		wg.Add(1)
		defer wg.Done()
		time.Sleep(time.Second * 20)
		fmt.Println("index")
	}
}

func SetupRouter(oriEngine *oriEngine.OriEngine) *gin.Engine {
	engine := gin.New()
	engine.Any("/", handle(oriEngine))
	return engine
}

func Run(oriEngine *oriEngine.OriEngine) {
	defer oriEngine.Wg.Done()
	flag.Parse()
	conf := config.GetHotConf()
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := SetupRouter(oriEngine)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Http.Port),
		Handler:        engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if *graceful {
		// 子进程监听父进程传递的 socket 描述符
		log.LogInfo("平滑重启-[子进程监听文件描述]")
		// 子进程的 0, 1, 2 是预留给标准输入、标准输出、错误输出
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		log.LogInfo("正常启动-[监听信号]")
		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		log.LogError("listener error: %+v", err)
		return
	}
	go func() {
		if err := server.Serve(listener); err != nil {
			log.LogInfo("http服务退出,err:%+v", err)
			return
		}
	}()
	signalHandle(oriEngine, server)
}

func signalHandle(oriEngine *oriEngine.OriEngine, server *http.Server) {
	for {
		select {
		case <-oriEngine.Context.Done():
			log.LogInfo("http服务退出")
			return
		case sig := <-oriEngine.HttpSignal:
			ctx, _ := context.WithCancel(context.Background())
			switch sig {
			case oriSignal.SIGUSR1: //优雅退出
				if err := server.Shutdown(ctx); err != nil {
					log.LogError(err.Error())
					break
				}
				log.LogInfo("http服务优雅退出")
				wg.Wait() //登陆用户连接全部断开
				oriEngine.Signal <- syscall.SIGINT
			case oriSignal.SIGUSR2: //平滑重启
				err := reload() // 执行热重启函数
				if err != nil {
					log.LogError("http服务reload error: %v", err)
					break
				}
				log.LogInfo("http服务热重启")
				if err := server.Shutdown(ctx); err != nil {
					log.LogError(err.Error())
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
	args := []string{"-graceful"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} // 文件描述符
	// 新建并执行子进程
	return cmd.Start()
}
