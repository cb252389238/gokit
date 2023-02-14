package http

import "github.com/go-resty/resty/v2"

func Get(url string) (string, error) {
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)
	return string(resp.Body()), err
}
