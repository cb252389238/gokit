package php

import (
	"io"
	"os"
	"path/filepath"
)

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
