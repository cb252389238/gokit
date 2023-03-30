//go:build linux
// +build linux

package oriSignal

import (
	"os"
	"syscall"
)

const (
	SIGUSR1 = syscall.SIGUSR1 //用户自定义信号1
	SIGUSR2 = syscall.SIGUSR2 //用户自定义信号2
)

// linux 信号
var Signals []os.Signal = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2}
