package tools

import (
	"net/url"
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

// 去除url保留path
func ParseReplaceUrl(urlPath string) string {
	parse, _ := url.Parse(urlPath)
	path := parse.Path
	if parse.RawQuery != "" {
		path += "?" + parse.RawQuery
	}
	if len(path) > 0 {
		if path[0] == '/' {
			path = path[1:]
		}
	}
	return path
}
