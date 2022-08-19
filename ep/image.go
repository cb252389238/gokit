package ep

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

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
