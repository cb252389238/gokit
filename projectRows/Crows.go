package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

var (
	SumRows   int64
	FilePaths []string
	l         sync.Mutex
	sufStr    string
	sufSlice  []string
)

func init() {
	flag.StringVar(&sufStr, "s", "", "-s 后缀名 例如 -s go,php")
}

func main() {
	t := time.Now()
	flag.Parse()
	if sufStr != "" {
		if strings.Contains(sufStr, ",") || strings.Contains(sufStr, "，") {
			sufStr := strings.ReplaceAll(sufStr, "，", ",")
			sufSlice = strings.Split(sufStr, ",")
		} else {
			sufSlice = append(sufSlice, sufStr)
		}
	}
	GetAllFileName("./")
	for _, v := range FilePaths {
		GetRowsNum(v)
	}
	since := time.Since(t)
	fmt.Println("代码总行数:", SumRows)
	fmt.Println("耗时:", since)
	for {
		time.Sleep(time.Second)
	}
}

func GetRowsNum(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		return
	}
	count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	fmt.Println(file)
	fmt.Println("|___rows:", count)
	l.Lock()
	SumRows += int64(count)
	l.Unlock()
}

func GetAllFileName(paths string) {
	files, err := os.ReadDir(paths)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range files {
		if v.IsDir() {
			GetAllFileName(paths + v.Name() + "/")
		} else {
			if sufStr != "" {
				if hasSuf(path.Ext(v.Name())) {
					FilePaths = append(FilePaths, paths+v.Name())
				}
			} else {
				FilePaths = append(FilePaths, paths+v.Name())
			}
		}
	}
}

func hasSuf(suf string) bool {
	suf = strings.ReplaceAll(suf, ".", "")
	isHas := false
	for _, v := range sufSlice {
		if v == suf {
			isHas = true
			break
		}
	}
	return isHas
}
