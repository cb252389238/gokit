package ep

import (
	"net/url"
)

func UrlEncode(urlstr string) string {
	encodeurl := url.QueryEscape(urlstr)
	return encodeurl
}

func UrlDecode(urlstr string) (string, error) {
	decodeurl, err := url.QueryUnescape(urlstr)
	if err != nil {
		return "", err
	} else {
		return decodeurl, nil
	}
}
