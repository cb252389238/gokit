package easy

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
	"time"
)

func MakeDir(path string) (string, error) {
	bools, _ := PathExists(path)
	if !bools {
		err := os.MkdirAll(path, 755)
		if err != nil {
			return "", err
		} else {
			return path, nil
		}
	} else {
		return "", nil
	}
}

// 返回全名，后缀名，文件名
func FileInfo(file string) (string, string, string) {
	fullName := path.Base(file)
	suffix := path.Ext(file)
	prefix := fullName[0 : len(fullName)-len(suffix)]
	return fullName,
		suffix,
		prefix
}

// 检查是否是图片文件
func CheckImageFile(path, style string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("打开文件失败 %s", err.Error())
	}
	switch strings.ToUpper(style) {
	case "JPG", "JPEG":
		_, err = jpeg.Decode(f)
	case "PNG":
		_, err = png.Decode(f)
	case "GIF":
		_, err = gif.Decode(f)
	}
	if err != nil {
		fmt.Errorf("校验文件类型失败 %s", err.Error())
		return "", err
	}
	return "", nil
}

// 写入文件
func WriteToFile(fileName, content, path string) error {
	fileName = fileName + ".logs"
	_, err := MakeDir(path)
	if err != nil {
		return err
	}
	fconn, err := MakeFile(path, fileName)
	if err != nil {
		return err
	}
	content = "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "\r\n" + content + "\r\n"
	_, err = fconn.WriteString(content)
	if err != nil {
		return err
	}
	defer fconn.Close()
	return nil
}

// 创建文件
func MakeFile(path, fileName string) (*os.File, error) {
	bools, _ := PathExists(path + "/" + fileName)
	if bools == false {
		fconn, err := os.Create(path + "/" + fileName)
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

// 判断文件或者文件夹是否存在
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
