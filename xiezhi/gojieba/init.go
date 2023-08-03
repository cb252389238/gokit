package gojieba

import (
	"xiezhi/gojieba/deps/cppjieba"
	"xiezhi/gojieba/deps/limonp"
	"xiezhi/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
