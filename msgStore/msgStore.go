package msgStore

import (
	"sync"
	"time"
)

const (
	ARR_LEN = 200 //存储数组长度
	MSG_LEN = 20  //返回消息长度
)

type Store struct {
	Bucket     *Cache
	incrIds    sync.Map
	lastOffset sync.Map
	l          sync.RWMutex
}

type message struct {
	context ImMsg
	ctime   int64
}

type ImMsg interface {
	GetRoomId() string
	GetUid() int
	GetData() map[string]interface{}
}

func NewMsgStore() *Store {
	return &Store{
		Bucket: NewCache(0, 0),
	}
}

func (s *Store) Set(value ImMsg) {
	s.l.Lock()
	defer s.l.Unlock()
	messageAny, b := s.Bucket.Get(value.GetRoomId())
	var offset int
	if !b {
		messageArr := &[ARR_LEN]message{}
		offset = s.getId(value.GetRoomId()) % ARR_LEN
		messageArr[offset] = message{
			context: value,
			ctime:   time.Now().Unix(),
		}
		s.Bucket.Set(value.GetRoomId(), messageArr, NoExpiration)
	} else {
		messageArr := messageAny.(*[ARR_LEN]message)
		offset = s.getId(value.GetRoomId()) % ARR_LEN
		messageArr[offset] = message{
			context: value,
			ctime:   time.Now().Unix(),
		}
	}
	s.setLastOffset(value.GetRoomId(), offset)
}

func (s *Store) Get(key string, uid, offset int) ([]map[string]interface{}, int) {
	s.l.RLock()
	defer s.l.RUnlock()
	messageAny, b := s.Bucket.Get(key)
	if !b {
		return []map[string]interface{}{}, offset
	}
	messageArr := messageAny.(*[ARR_LEN]message)
	data := []map[string]interface{}{}
	for i := 0; i < MSG_LEN; i++ {
		last := offset - 1
		if offset == 0 {
			last = ARR_LEN - 1
		}
		//历史得消息不返回
		if messageArr[offset].ctime < messageArr[last].ctime {
			break
		}
		//超过十秒前得消息不再返回
		if time.Now().Unix()-messageArr[offset].ctime > 10 {
			offset = s.getLastOffset(key)
			break
		}
		//自己的消息过滤
		if messageArr[offset].context.GetUid() == uid {
			continue
		}
		data = append(data, messageArr[offset].context.GetData())
		if offset >= ARR_LEN-1 {
			offset = 0
		} else {
			offset++
		}
	}
	return data, offset
}

func (s *Store) getId(key string) int {
	loadVal, ok := s.incrIds.Load(key)
	if !ok {
		s.incrIds.Store(key, 0)
		return 0
	}
	val := loadVal.(int) + 1
	s.incrIds.Store(key, val)
	return val
}

func (s *Store) getLastOffset(key string) int {
	loadVal, ok := s.lastOffset.Load(key)
	if !ok {
		s.lastOffset.Store(key, 0)
		return 0
	}
	return loadVal.(int)
}

func (s *Store) setLastOffset(key string, offset int) {
	s.lastOffset.Store(key, offset)
}
