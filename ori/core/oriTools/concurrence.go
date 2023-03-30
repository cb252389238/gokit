package oriTools

import (
	"sync"
)

var (
	concurrencyLock = sync.Mutex{}
	ConcurrencyNum  sync.Map
)

func ConcurrencyAdd(name string) {
	concurrencyLock.Lock()
	defer concurrencyLock.Unlock()
	value, ok := ConcurrencyNum.Load(name)
	if !ok {
		ConcurrencyNum.Store(name, 1)
	} else {
		ConcurrencyNum.Store(name, value.(int)+1)
	}
	value, ok = ConcurrencyNum.Load("all")
	if !ok {
		ConcurrencyNum.Store("all", 1)
	} else {
		ConcurrencyNum.Store("all", value.(int)+1)
	}
}

func ConcurrencyDel(name string) {
	concurrencyLock.Lock()
	defer concurrencyLock.Unlock()
	value, ok := ConcurrencyNum.Load(name)
	if ok {
		ConcurrencyNum.Store(name, value.(int)-1)
	}
	value, ok = ConcurrencyNum.Load("all")
	if ok {
		ConcurrencyNum.Store("all", value.(int)-1)
	}
}
