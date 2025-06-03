package tools

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// 获取当前程序运行路径
func getAbsBinPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst + "/"
}

// 获取当前用户路径
func getUserBinPath() string {
	var cpath string = ""
	path, _ := os.Getwd()
	cpath = path + "/"
	return cpath
}

// 获取项目根目录
func GetRootPath() string {
	if runtime.GOOS == "linux" {
		return getAbsBinPath()
	} else {
		return getUserBinPath()
	}
}
