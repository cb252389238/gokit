package php

import (
	"fmt"
	"os"
	"testing"
)

func TestFile_exists(t *testing.T) {
	f := "./array.go"
	fmt.Println(File_exists(f))
	fmt.Println(File_exists("./aa.go"))
}

func TestIs_file(t *testing.T) {
	fmt.Println(Is_file("./file.go"))
	fmt.Println(Is_file("../php"))
}

func TestIs_dir(t *testing.T) {
	fmt.Println(Is_dir("../php"))
	fmt.Println(Is_dir("./net.go"))
}

func TestFilesize(t *testing.T) {
	fmt.Println(Filesize("./file.go"))
}

func TestFile_get_contents(t *testing.T) {
	File_put_contents("./filetest.txt", "222", os.ModeAppend)
	fmt.Println(File_get_contents("./filetest.txt"))
}

func TestDelete(t *testing.T) {
	Delete("./filetest.txt")
}
