package varChannel

import (
	"errors"
	"time"
)

type varChannel struct {
	c     chan interface{} //读取通道
	list  *List            //写入链表
	len   int              //长度
	close bool             //是否关闭通道
}

func New() *varChannel {
	v := &varChannel{
		c:    make(chan interface{}, 50),
		list: NewList(),
		len:  0,
	}
	go func(v *varChannel) {
		for {
			if v.len > 0 {
				front := v.list.Front()
				if value := front.Value; value != nil {
					v.c <- value
					v.list.Remove(front)
					v.len--
				}
				continue
			}
			if v.close {
				close(v.c)
				return
			}
			time.Sleep(time.Millisecond * 1)
		}
	}(v)
	return v
}

func (v *varChannel) Write(value interface{}) error {
	if v.close {
		return errors.New("varChannel is closed")
	}
	v.list.PushBack(value)
	v.len++
	return nil
}

func (v *varChannel) Read() <-chan interface{} {
	return v.c
}

func (v *varChannel) Close() {
	v.close = true
}
