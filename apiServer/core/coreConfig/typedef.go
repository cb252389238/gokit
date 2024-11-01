package coreConfig

import (
	"apiServer/typedef"
	"sync"
)

type HotConf struct {
	Conf           typedef.Config
	L              sync.RWMutex
	LastModifyTime int64
}
