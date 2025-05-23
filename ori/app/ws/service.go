package ws

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"net"
	"net/http"
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
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	graceful              = flag.Bool("graceful-ws", false, "listen on fd open 3 (internal use only)")
	listener net.Listener = nil
	err      error
	wg       = &sync.WaitGroup{}
)

func handle(oriEngine *engine.OriEngine) gin.HandlerFunc {
	return func(c *gin.Context) {
		wg.Add(1)
		defer wg.Done()
		//升级get请求为webSocket协议
		_, err := upGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			oriLog.Error("%+v", err)
			return
		}
	}
}

func SetupRouter(oriEngine *engine.OriEngine) *gin.Engine {
	engine := gin.New()
	engine.GET("/*paramsUrl", handle(oriEngine))
	return engine
}

func Run(oriEngine *engine.OriEngine) {
	defer oriEngine.Wg.Done()
	if !flag.Parsed() {
		flag.Parse()
	}
	conf := oriConfig.GetHotConf()
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := SetupRouter(oriEngine)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", oriConfig.GetHotConf().Websocket.Port),
		Handler:        engine,
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
			oriLog.Info("websocket服务退出,err:%+v", err)
			return
		}
	}()
	signalHandle(oriEngine, server)
}

func signalHandle(oriEngine *engine.OriEngine, server *http.Server) {
	for {
		select {
		case <-oriEngine.Context.Done():
			oriLog.Info("websocket服务退出")
			return
		case sig := <-oriEngine.WsSignal:
			ctx, _ := context.WithCancel(context.Background())
			switch sig {
			case oriSignal.SIGUSR1: //优雅退出
				if err := server.Shutdown(ctx); err != nil {
					oriLog.Error(err.Error())
					break
				}
				oriLog.Info("websocket服务优雅退出")
				wg.Wait() //登陆用户连接全部断开
				oriEngine.Signal <- syscall.SIGINT
			case oriSignal.SIGUSR2: //平滑重启
				err := reload() // 执行热重启函数
				if err != nil {
					oriLog.Error("websocket服务reload error: %v", err)
					break
				}
				oriLog.Info("websocket服务热重启")
				if err := server.Shutdown(ctx); err != nil {
					oriLog.Error(err.Error())
					break
				}
				wg.Wait() //登陆用户连接全部断开
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
	pid := time.Now().Unix()
	args := []string{"-graceful-ws", cast.ToString(pid)}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout         // 标准输出
	cmd.Stderr = os.Stderr         // 错误输出
	cmd.ExtraFiles = []*os.File{f} // 文件描述符
	// 新建并执行子进程
	return cmd.Start()
}
