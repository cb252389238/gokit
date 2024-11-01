package corePool

import "errors"

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("corePool is closed")
)

// Pool 基本方法
type Pool interface {
	Get() (interface{}, error)

	Put(interface{}) error

	Close(interface{}) error

	Release()

	Len() int
}
