package ws

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"frame/internal/svc"
	"frame/resource"
	"frame/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	graceful              = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
	listener net.Listener = nil
	err      error
	wg       = &sync.WaitGroup{}
)

func handle(ctx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		wg.Add(1)
		defer wg.Done()
		//升级get请求为webSocket协议
		_, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			resource.LogError("%+v", err)
			return
		}
	}
}

func SetupRouter(ctx *svc.ServiceContext) *gin.Engine {
	engine := gin.New()
	engine.GET("/*paramsUrl", handle(ctx))
	return engine
}

func Run(ctx *svc.ServiceContext) {
	defer ctx.Wg.Done()
	flag.Parse()
	conf := resource.GetAllHotConf()
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := SetupRouter(ctx)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", resource.GetAllHotConf().Websocket.Port),
		Handler:        engine,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if *graceful {
		// 子进程监听父进程传递的 socket 描述符
		resource.LogInfo("平滑重启-[子进程监听文件描述]")
		// 子进程的 0, 1, 2 是预留给标准输入、标准输出、错误输出
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		resource.LogInfo("正常启动-[监听信号]")
		listener, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		resource.LogError("listener error: %+v", err)
		return
	}
	go func() {
		if err := server.Serve(listener); err != nil {
			resource.LogInfo("websocket协程退出,err:%+v", err)
			ctx.C <- syscall.SIGHUP
			return
		}
	}()
	signalHandle(ctx, server)
}

func signalHandle(ctx *svc.ServiceContext, server *http.Server) {
	for {
		select {
		case <-ctx.Context.Done():
			resource.LogInfo("ws协程退出")
			return
		case sig := <-ctx.WsC:
			ctx2, _ := context.WithCancel(context.Background())
			switch sig {
			case types.SIGUSR1: //优雅退出
				if err := server.Shutdown(ctx2); err != nil {
					resource.LogError(err.Error())
					break
				}
				resource.LogInfo("优雅退出程序")
				wg.Wait() //登陆用户连接全部断开
				ctx.C <- syscall.SIGINT
			case types.SIGUSR2: //平滑重启
				err := reload() // 执行热重启函数
				if err != nil {
					resource.LogError("reload error: %v", err)
					break
				}
				if err := server.Shutdown(ctx2); err != nil {
					resource.LogError(err.Error())
					break
				}
				wg.Wait() //登陆用户连接全部断开
				ctx.C <- syscall.SIGINT
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
	pid := time.Now().Unix()
	args := []string{"-graceful", cast.ToString(pid)}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} // 文件描述符
	// 新建并执行子进程
	return cmd.Start()
}
