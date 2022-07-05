package file

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadFile(url string, params map[string]string, nameField, fileName string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	HttpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

//文件中换二进制
func File2Bytes(paths string) ([]byte, error) {
	file, err := os.Open(paths)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Copy(srcFile, dstFile string) error {
	src, err := os.OpenFile(srcFile, os.O_RDONLY, 0666)
	if err != nil {
		return errors.New("os.OpenFile is error")
	}
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return errors.New("os.OpenFile is error")
	}
	defer func() {
		src.Close()
		dst.Close()
	}()
	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dst)
	buffer := make([]byte, 1024)
	for {
		_, err = reader.Read(buffer)
		if err != nil && err != io.EOF {
			return errors.New("Read to buffer error")
		} else {
			if err == io.EOF {
				return nil
			}
			_, err := writer.Write(buffer)
			if err != nil {
				return errors.New("Write to buffer error")
			}
		}
	}
}
