package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	IMG_DIR_TRUE  string = "./img/"
	IMG_DIR_FALSE string = "/images/"
	HTTP_PORT     string = ":9001"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

var (
	CallFuncs = map[string]HandlerFunc{
		"index":  Index,
		"images": Images,
		"login":  Login,
	}
	Methods = []string{"index", "images", "login"}
)

type HttpHander struct {
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", Login)
	http.HandleFunc(IMG_DIR_FALSE, Images)
	err := http.ListenAndServe(HTTP_PORT, HttpHander{})
	if err != nil {
		fmt.Println(err)
	}
}

func (this HttpHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path
	uri = strings.Trim(uri, "/")
	if uri == "" {
		Index(w, r)
		return
	}
	sclie_uri := strings.Split(uri, "/")
	if len(sclie_uri) != 0 {
		uri = sclie_uri[0]
	}
	hasN := 0
	for _, v := range Methods {
		if v == uri {
			hasN++
		}
	}
	if hasN == 0 {
		Index(w, r)
		return
	}
	CallFuncs[uri](w, r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login")
}

func Images(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[len(IMG_DIR_FALSE):] == "" {
		fmt.Fprintln(w, "错路的路径")
		return
	}
	file := IMG_DIR_TRUE + r.URL.Path[len(IMG_DIR_FALSE):]
	f, err := os.Open(file)
	if err != nil && os.IsNotExist(err) {
		file = IMG_DIR_TRUE + "/1.jpg"
	}
	defer f.Close()
	http.ServeFile(w, r, file)
	return
}
