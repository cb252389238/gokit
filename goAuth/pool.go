package coreAuth

import "errors"

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)

// Pool 基本方法
type Pool interface {
	Get() (any, error)

	Put(any) error

	Close(any) error

	Release()

	Len() int
}
