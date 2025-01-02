package imMerge

import "sync"

const (
	copyThreshold = 1000  //复制阈值
	maxDeletion   = 10000 //最大删除数量
)

type SafeMap struct {
	lock        sync.RWMutex //锁
	deletionOld int
	deletionNew int
	dirtyOld    map[any]any
	dirtyNew    map[any]any
}

// 初始化SafeMap
func NewSafeMap() *SafeMap {
	return &SafeMap{
		dirtyOld: make(map[any]any),
		dirtyNew: make(map[any]any),
	}
}

// Del deletes the value with the given key from m.
func (m *SafeMap) Del(key any) {
	m.lock.Lock()
	if _, ok := m.dirtyOld[key]; ok {
		delete(m.dirtyOld, key)
		m.deletionOld++
	} else if _, ok := m.dirtyNew[key]; ok {
		delete(m.dirtyNew, key)
		m.deletionNew++
	}
	if m.deletionOld >= maxDeletion && len(m.dirtyOld) < copyThreshold {
		for k, v := range m.dirtyOld {
			m.dirtyNew[k] = v
		}
		m.dirtyOld = m.dirtyNew
		m.deletionOld = m.deletionNew
		m.dirtyNew = make(map[any]any)
		m.deletionNew = 0
	}
	if m.deletionNew >= maxDeletion && len(m.dirtyNew) < copyThreshold {
		for k, v := range m.dirtyNew {
			m.dirtyOld[k] = v
		}
		m.dirtyNew = make(map[any]any)
		m.deletionNew = 0
	}
	m.lock.Unlock()
}

// 获取值
func (m *SafeMap) Get(key any) (any, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.dirtyOld[key]; ok {
		return val, true
	}
	val, ok := m.dirtyNew[key]
	return val, ok
}

// 设置值
func (m *SafeMap) Set(key, value any) {
	m.lock.Lock()
	if m.deletionOld <= maxDeletion {
		if _, ok := m.dirtyNew[key]; ok {
			delete(m.dirtyNew, key)
			m.deletionNew++
		}
		m.dirtyOld[key] = value
	} else {
		if _, ok := m.dirtyOld[key]; ok {
			delete(m.dirtyOld, key)
			m.deletionOld++
		}
		m.dirtyNew[key] = value
	}
	m.lock.Unlock()
}

// 返回元素数量
func (m *SafeMap) Size() int {
	m.lock.RLock()
	size := len(m.dirtyOld) + len(m.dirtyNew)
	m.lock.RUnlock()
	return size
}
