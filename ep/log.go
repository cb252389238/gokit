package ep

import (
	"os"
	"sync"
	"time"
)

var mux sync.RWMutex
func WriteLog(fileName,content,path string) error {
	mux.Lock()	//进行读写锁
	fileName = fileName+".log"
	path,err := MakeDir(path)
	if err != nil {
		return err
	}
	fconn, err := MakeFile(path,fileName)
	if err != nil {
		return err
	}
	content = "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"\r\n"+content+"\r\n"
	_, err = fconn.WriteString(content)
	if err != nil {
		return err
	}
	defer mux.Unlock()
	defer fconn.Close()
	return nil
}

//创建目录目录根据当天日期进行创建
func MakeDir(path string) (string,error) {
	err := os.MkdirAll(path, 775)
	if err != nil {
		return "",err
	} else {
		return path,nil
	}
}

//创建文件用来记录日志
func MakeFile(path,fileName string) (*os.File, error) {
	bools, _ := PathExists(path+"/"+fileName)
	if bools == false {
		fconn, err := os.Create(path+"/"+fileName)
		if err != nil {
			return fconn, err
		} else {
			return fconn, nil
		}
	} else {
		fconn, err := os.OpenFile(path+"/"+fileName, os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			return fconn, err
		} else {
			return fconn, err
		}
	}
}

//判断文件或者文件夹是否存在
func PathExists(paths string) (bool, error) {
	_, err := os.Stat(paths)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
