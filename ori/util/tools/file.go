package tools

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	var size int64 = 1048576 // 1M
	hash := md5.New()

	if fi.Size() < size {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		hash.Write(data)
	} else {
		b := make([]byte, size)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}

			hash.Write(b[:n])
		}
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
