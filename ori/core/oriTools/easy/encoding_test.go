package easy

import "testing"

func TestConvert(t *testing.T) {
	dstCharset := "gbk"
	srcCharset := "utf-8"
	src := "你好"
	dst, err := Convert(dstCharset, srcCharset, src)
	if err != nil {
		t.Error(err)
	}
	t.Log(dst)
	t.Log(ToUTF8(dstCharset, dst))
	t.Log(UTF8To(srcCharset, dst))
}

func TestUrlEncode(t *testing.T) {
	urlStr := "http://www.baidu.com/?a=1&b=2"
	t.Log(UrlEncode(urlStr))
	t.Log(UrlDecode(UrlEncode(urlStr)))
}
