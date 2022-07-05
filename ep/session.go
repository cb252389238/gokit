package ep

import (
	"crypto/rand"
	"fmt"
	"io"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

type Session struct {
	L   sync.Mutex
	Age int64 //寿命 秒
}

type ValueStr struct {
	Data interface{}
	Age  int64
}

func (this *Session) SessionId() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func (this *Session) SetSession(value interface{}) (string, error) {
	SessionId, err := this.SessionId()
	if err != nil {
		return "", err
	}
	var data ValueStr
	data = ValueStr{Data: value, Age: time.Now().Unix() + this.Age}
	sessionMap.Store(SessionId, data)
	return SessionId, nil
}

func (this *Session) GetSession(SessionId interface{}) (interface{}, bool) {
	value, ok := sessionMap.Load(SessionId)
	if !ok {
		return "", false
	}
	if valueT, ok := value.(ValueStr); ok {
		if valueT.Age <= time.Now().Unix() {
			sessionMap.Delete(SessionId)
			return "", false
		}
		return valueT, true
	} else {
		return "", false
	}
}
