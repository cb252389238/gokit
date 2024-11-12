package easy

import (
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// 创建目录
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

// 返回文件真实类型
func FileType(filePath string) (string, error) {
	var fileType string
	var err error
	file, err := os.Open(filePath)
	if err != nil {
		return fileType, err
	}
	defer file.Close()
	// 读取文件的前 512 个字节
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if n <= 0 {
		return fileType, errors.New("empty file")
	}
	if err != nil {
		return fileType, err
	}
	// 调用 http.DetectContentType 方法判断文件类型
	contentType := http.DetectContentType(buffer[:n])
	byType, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return fileType, err
	}
	return byType[len(byType)-1][1:], nil
}

func File_exists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func Is_file(filename string) bool {
	fd, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return !fd.IsDir()
}

func Is_dir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

func Filesize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

func File_put_contents(filename string, data string, mode os.FileMode) error {
	return os.WriteFile(filename, []byte(data), mode)
}

func File_get_contents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

func Delete(filename string) error {
	return os.Remove(filename)
}

func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

func Basename(path string) string {
	return filepath.Base(path)
}

func Fclose(handle *os.File) error {
	return handle.Close()
}
