package easy

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
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

func PostJson(url string, data []byte) (string, error) {
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

func PostForm(url string, data url.Values) (string, error) {
	resp, err := http.PostForm(url, data)
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
