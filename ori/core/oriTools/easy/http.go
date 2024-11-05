package easy

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Post(url string, data []byte) (string, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}

func HttpDo(httpMethod string, url string, headerMap map[string]string, rawBody string) (remoteResponse string, err error) {
	client := &http.Client{}
	req, err0 := http.NewRequest(httpMethod, url, strings.NewReader(rawBody))
	if err0 != nil {
		err = err0
		return
	}
	if len(headerMap) > 0 {
		for k, v := range headerMap {
			req.Header.Set(k, v)
		}
	}
	resp, err1 := client.Do(req)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}
	remoteResponse = string(body)
	return
}
