package http

import (
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//func main() {
//	res, _ := HttpDo("POST", "http://localhost/test/20190228/index.php", "id=123&name=chengbang", map[string]string{"session_id": "123", "token": "456"})
//	fmt.Println(res)
//}

func HttpGet(urlStr string) (string, error) {

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPost(urlStr string, data string) (string, error) {
	//	data = "name=chengbang&age=28&sex=男"
	resp, err := http.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPostForm(urlStr string, data map[string]string) (string, error) {
	postData := url.Values{}
	for k, v := range data {
		postData.Set(k, v)
	}
	resp, err := http.PostForm(urlStr, postData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpDo(method string, urlStr string, data string, cookies map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data))
	if err != nil {
		return "", err
	}
	//设置header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//设置cookie
	for k, v := range cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v, HttpOnly: true})
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
