package rest

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getAbsBinPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst + "/"
}

func getUserBinPath() string {
	var cpath string = ""
	path, _ := os.Getwd()
	cpath = path + "/"
	return cpath
}

func GetRootPath() string {
	if runtime.GOOS == "linux" {
		return getAbsBinPath()
	} else {
		return getUserBinPath()
	}
}
