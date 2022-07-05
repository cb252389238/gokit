package ep

import (
	"os"
	"os/exec"
	"path/filepath"
)

/**
获取用户执行目录
*/
func GetUserBinPath() string {
	path, _ := os.Getwd()
	return path + "/"
}

/**
获取程序执行绝对路径
*/
func GetAbsBinPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst + "/"
}

/**
获取执行程序相对路径
*/

func GetRelBinPath() string {
	path, _ := exec.LookPath(os.Args[0])
	path = filepath.Dir(path)
	return "./" + path + "/"
}
