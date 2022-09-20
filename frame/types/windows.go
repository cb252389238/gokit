//go:build windows
// +build windows

package types

import (
	"os"
	"syscall"
)

const (
	SIGUSR1 = syscall.SIGHUP  //用户自定义信号1
	SIGUSR2 = syscall.SIGQUIT //用户自定义信号2
)

//windows信号
var Signals []os.Signal = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
