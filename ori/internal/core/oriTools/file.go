package oriTools

import (
	"os"
	"path"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MakeDir(path string) (string, error) {
	bools, _ := pathExists(path)
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
