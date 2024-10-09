package helper

import (
	"github.com/gorilla/websocket"
	"net/url"
)

// 检测im服务器是否可用
func CheckWebSocketConnection(urlStr string) (bool, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false, err
	}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}
