package http

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	get, err := HttpGet("http://127.0.0.1:9002/ad/baiduAttrCallback?uid=1")
	fmt.Println(get, err)
}

func TestHttpPost(t *testing.T) {
	get, err := HttpPost("http://127.0.0.1:9002/ad/baiduAttrCallback?uid=1", `{"uid":1001}`)
	fmt.Println(get, err)
}
