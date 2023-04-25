package easy

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// 获取程序执行绝对路径
func getAbsBinPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst + "/"
}

// 获取用户执行目录
func getUserBinPath() string {
	var cpath string = ""
	path, _ := os.Getwd()
	cpath = path + "/"
	return cpath
}

// 获取执行程序相对路径
func GetRelBinPath() string {
	path, _ := exec.LookPath(os.Args[0])
	path = filepath.Dir(path)
	return "./" + path + "/"
}

func GetRootPath() string {
	if runtime.GOOS == "linux" {
		return getAbsBinPath()
	} else {
		return getUserBinPath()
	}
}
