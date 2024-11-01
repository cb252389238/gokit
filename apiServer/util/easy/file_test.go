package easy

import (
	"fmt"
	"testing"
)

func TestFileType(t *testing.T) {
	fileType, err := FileType("1.jpg")
	fmt.Println(fileType, err)
}
