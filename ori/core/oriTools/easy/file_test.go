package easy

import "testing"

func TestMakeDir(t *testing.T) {
	path := "D:/test/test"
	res, err := MakeDir(path)
	t.Log(res, err)
}

func TestFileInfo(t *testing.T) {
	path := "D:/test/test/test.txt"
	name, size, name2 := FileInfo(path)
	t.Log(name, size, name2)
}

func TestCheckImageFile(t *testing.T) {
	path := "D:/test/test/123.txt"
	ok := CheckImageFile(path, "png")
	t.Log(ok)
	path2 := "D:/test/test/boy.png"
	ok = CheckImageFile(path2, "png")
	t.Log(ok)
}

func TestMakeFile(t *testing.T) {
	path := "D:/test/test/"
	file, err := MakeFile(path, "test.txt")
	t.Log(file, err)
}

func TestPathExists(t *testing.T) {
	path := "D:/test/test/"
	ok, err := PathExists(path)
	t.Log(ok, err)
	path2 := "D:/test/test/test.txt"
	ok, err = PathExists(path2)
	t.Log(ok, err)
	path3 := "D:/test/test/test2.txt"
	ok, err = PathExists(path3)
	t.Log(ok, err)
}

func TestFileType(t *testing.T) {
	path := "D:/test/test/test.txt"
	fileType, err := FileType(path)
	t.Log(fileType, err)
	path2 := "D:/test/test/boy.png"
	fileType, err = FileType(path2)
	t.Log(fileType, err)
	path3 := "D:/test/test/boy.jpg"
	fileType, err = FileType(path3)
}

func TestFileExists(t *testing.T) {
	path := "D:/test/test/test.txt"
	ok := FileExists(path)
	t.Log(ok)
	path2 := "D:/test/test/boy.png"
	ok = FileExists(path2)
	t.Log(ok)
	path3 := "D:/test/test/boy.jpg"
	ok = FileExists(path3)
	t.Log(ok)
}

func TestIsFile(t *testing.T) {
	path := "D:/test/test/test.txt"
	ok := IsFile(path)
	t.Log(ok)
	path2 := "D:/test/test/boy.png"
	ok = IsFile(path2)
	t.Log(ok)
	path3 := "D:/test/test"
	ok = IsFile(path3)
	t.Log(ok)
}

func TestIsDir(t *testing.T) {
	path := "D:/test/test/test.txt"
	ok, err := IsDir(path)
	t.Log(ok, err)
	path2 := "D:/test/test/boy.png"
	ok, err = IsDir(path2)
	t.Log(ok, err)
	path3 := "D:/test/test"
	ok, err = IsDir(path3)
	t.Log(ok, err)
}

func TestFilesize(t *testing.T) {
	path := "D:/test/test/test.txt"
	filesize, err := Filesize(path)
	t.Log(filesize, err)
	path2 := "D:/test/test/boy.png"
	filesize, err = Filesize(path2)
	t.Log(filesize, err)
	path3 := "D:/test/test/boy.jpg"
	filesize, err = Filesize(path3)
	t.Log(filesize, err)
}

func TestDelete(t *testing.T) {
	path := "D:/test/test/123.txt"
	err := Delete(path)
	t.Log(err)
}

func TestCopy(t *testing.T) {
	path := "D:/test/test/test.txt"
	path2 := "D:/test/test/test2.txt"
	ok, err := Copy(path, path2)
	t.Log(ok, err)
}

func TestRename(t *testing.T) {
	path := "D:/test/test/test2.txt"
	path2 := "D:/test/test/test3.txt"
	err := Rename(path, path2)
	t.Log(err)
}

func TestRealpath(t *testing.T) {
	path := "D:/test/test/test.txt"
	realpath, err := AbsPath(path)
	t.Log(realpath, err)
	path2 := "D:/test/test/boy.png"
	realpath = Basename(path2)
	t.Log(realpath)
}

func TestAppendFile(t *testing.T) {
	path := "D:/test/test/test.txt"
	err := AppendFile(path, "hello world")
	t.Log(err)
}

func TestOverwriteFile(t *testing.T) {
	path := "D:/test/test/test.txt"
	err := OverwriteFile(path, "hello world")
	t.Log(err)
}
