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
func CheckImageFile(path, style string) bool {
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
	default:
		return false
	}
	if err != nil {
		return false
	}
	return true
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

// 判断文件是否存在
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断是否文件
func IsFile(filename string) bool {
	fd, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return !fd.IsDir()
}

// 判断是否目录
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// 文件大小
func Filesize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// 读取文件内容
func FileGetContents(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	return string(data), err
}

// 删除文件
func Delete(filename string) error {
	return os.Remove(filename)
}

// 复制文件
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

// 更改文件名
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// 返回给定路径的绝对路径。
func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}

// 返回文件名
func Basename(path string) string {
	return filepath.Base(path)
}

// 向文件内追加内容
func AppendFile(filePath string, content string) error {
	// 以追加模式打开文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 向文件尾部追加内容
	_, err = file.WriteString(content + "\n")
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	return nil
}

// 覆盖文件内容
func OverwriteFile(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0666)
}
