package speed

import (
	"context"
	"sync"
	"time"
)

type Cache struct {
	*cache
}

type cache struct {
	kvItems        map[string]KVItem   //k-v结构
	kv_mu          sync.RWMutex        //读写锁
	hashItems      map[string]HASHItem //hash结构
	hash_mu        sync.RWMutex
	setItems       map[string]SetItem //集合
	set_mu         sync.RWMutex
	deleteCallBack func(string, any) //回调事件  超时或者删除的时候触发回调
	snowflake      *Node             //雪花算法生成key
	timeWheel      *TimeWheel        //时间轮  过期调用
	ctx            context.Context
	cancel         context.CancelFunc
}

type KVItem struct {
	Object     any   //存储体
	Expiration int64 //过期时间
	CallBack   bool  //是否回调
	Key        string
}

type HASHItem struct {
	Object     map[string]any //存储体
	Expiration int64          //过期时间
	CallBack   bool           //是否回调
	Key        string
}

type SetItem struct {
	Object map[any]Set //存储体
}

type Set struct {
	Key          string
	timeWheelKey string
	Member       any
	Expiration   int64 //过期时间
	CallBack     bool  //是否回调
}

func New() (*Cache, error) {
	ip := GetLoaclIp()
	node := Ipv4StringToInt(ip) % 256
	sf, err := NewNode(node)
	if err != nil {
		return nil, err
	}
	tw := NewTw(time.Second, 60, nil)
	tw.Start()
	ctx, cancelFunc := context.WithCancel(context.Background())
	c := &Cache{&cache{
		kvItems:        map[string]KVItem{},
		hashItems:      map[string]HASHItem{},
		setItems:       map[string]SetItem{},
		deleteCallBack: nil,
		snowflake:      sf,
		timeWheel:      tw,
		ctx:            ctx,
		cancel:         cancelFunc,
	}}
	go c.run()
	return c, nil
}

func (c *cache) run() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case data := <-c.timeWheel.C: //超时队列
			switch v := data.(type) {
			case KVItem:
				c.kv_mu.Lock()
				if i, b := c.kvDelete(v.Key); b {
					if v.CallBack {
						c.deleteCallBack(i.Key, i.Object)
					}
				}
				c.kv_mu.Unlock()
			case HASHItem:
				c.hash_mu.Lock()
				if i, b := c.hashDelete(v.Key); b {
					if v.CallBack {
						c.deleteCallBack(i.Key, i.Object)
					}
				}
				c.hash_mu.Unlock()
			case Set:
				c.set_mu.Lock()
				if i, b := c.setDelete(v.Key, v.Member); b {
					if v.CallBack {
						c.deleteCallBack(i.Key, i.Member)
					}
				}
				c.set_mu.Unlock()
			}
		}
	}
}

func (c *cache) BindDeleteCallBackFunc(f func(string, any)) {
	c.kv_mu.Lock()
	c.deleteCallBack = f
	c.kv_mu.Unlock()
}

func (c *cache) Stop() {
	c.cancel()
}

func (c *cache) Set(k string, v any, d time.Duration, callBack bool) {
	var endTime int64
	if d > 0 {
		endTime = time.Now().Add(d).Unix()
	}
	c.kv_mu.RLock()
	val, ok := c.kvItems[k]
	c.kv_mu.RUnlock()
	if ok {
		if val.Expiration > 0 {
			c.timeWheel.RemoveTimer(k)
		}
	}
	item := KVItem{
		Object:     v,
		Expiration: endTime,
		CallBack:   callBack,
		Key:        k,
	}
	c.kv_mu.Lock()
	c.kvItems[k] = item
	c.kv_mu.Unlock()
	c.timeWheel.AddTimer(d, k, item)
}

func (c *cache) SetNx(k string, v any, d time.Duration, callBack bool) bool {
	var endTime int64
	if d > 0 {
		endTime = time.Now().Add(d).Unix()
	}
	c.kv_mu.Lock()
	defer c.kv_mu.Unlock()
	_, ok := c.kvItems[k]
	if ok {
		return false
	}
	item := KVItem{
		Object:     v,
		Expiration: endTime,
		CallBack:   callBack,
		Key:        k,
	}
	c.kvItems[k] = item
	c.timeWheel.AddTimer(d, k, item)
	return true
}

func (c *cache) Get(k string) (any, bool) {
	c.kv_mu.RLock()
	item, ok := c.kvItems[k]
	c.kv_mu.RUnlock()
	if !ok {
		return nil, false
	}
	if item.Expiration <= time.Now().Unix() {
		return nil, false
	}
	return item.Object, true
}

// 获取k-v 过期时间
func (c *cache) GetEx(k string) (any, time.Time, bool) {
	c.kv_mu.RLock()
	item, ok := c.kvItems[k]
	c.kv_mu.RUnlock()
	if !ok {
		return nil, time.Time{}, false
	}
	if item.Expiration <= time.Now().Unix() {
		return nil, time.Time{}, false
	}
	return item.Object, time.Unix(item.Expiration, 0), true
}

// k-v删除
func (c *cache) Del(k string) {
	c.kv_mu.Lock()
	v, ok := c.kvDelete(k)
	c.kv_mu.Unlock()
	if v.Expiration > 0 && ok {
		c.timeWheel.RemoveTimer(k)
	}
	if ok && v.CallBack && c.deleteCallBack != nil {
		c.deleteCallBack(v.Key, v.Object)
	}
}

func (c *cache) kvDelete(k string) (KVItem, bool) {
	if v, ok := c.kvItems[k]; ok {
		delete(c.kvItems, k)
		return v, true
	}
	return KVItem{}, false
}

// 获取k-v所有值
func (c *cache) Items() map[string]any {
	c.kv_mu.RLock()
	defer c.kv_mu.RUnlock()
	m := make(map[string]any, len(c.kvItems))
	now := time.Now().Unix()
	for k, v := range c.kvItems {
		if now > v.Expiration {
			continue
		}
		m[k] = v
	}
	return m
}

// 获取k-v数量
func (c *cache) ItemCount() int {
	c.kv_mu.RLock()
	n := len(c.kvItems)
	c.kv_mu.RUnlock()
	return n
}

// 判断k-v值是否存在
func (c *cache) Exists(k string) bool {
	c.kv_mu.RLock()
	_, ok := c.kvItems[k]
	c.kv_mu.RUnlock()
	return ok
}

func (c *cache) hashDelete(k string) (HASHItem, bool) {
	if v, ok := c.hashItems[k]; ok {
		delete(c.hashItems, k)
		return v, true
	}
	return HASHItem{}, false
}

func (c *cache) HSet(key, field string, val any) {
	c.hash_mu.RLock()
	hash, ok := c.hashItems[key]
	c.hash_mu.RUnlock()
	if ok {
		c.hash_mu.Lock()
		hash.Object[field] = val
		c.hash_mu.Unlock()
		return
	}
	item := HASHItem{
		Object: map[string]any{field: val},
		Key:    key,
	}
	c.hash_mu.Lock()
	c.hashItems[key] = item
	c.hash_mu.Unlock()
}

func (c *cache) HSetEx(key string, d time.Duration, callBack bool) bool {
	var endTime int64
	if d > 0 {
		endTime = time.Now().Add(d).Unix()
	}
	c.hash_mu.RLock()
	hash, ok := c.hashItems[key]
	c.hash_mu.RUnlock()
	if ok {
		if hash.Expiration > 0 {
			c.timeWheel.RemoveTimer(key)
		}
		c.hash_mu.Lock()
		hash.CallBack = callBack
		hash.Expiration = endTime
		c.hash_mu.Unlock()
		c.timeWheel.AddTimer(d, key, hash)
		return true
	}
	return false
}

func (c *cache) HMSet(key string, data map[string]any) {
	if len(data) == 0 {
		return
	}
	c.hash_mu.RLock()
	hash, ok := c.hashItems[key]
	c.hash_mu.RUnlock()
	if ok {
		c.hash_mu.Lock()
		for field, value := range data {
			hash.Object[field] = value
		}
		c.hash_mu.Unlock()
		return
	}
	item := HASHItem{
		Key: key,
	}
	for field, value := range data {
		item.Object[field] = value
	}
	c.hash_mu.Lock()
	c.hashItems[key] = item
	c.hash_mu.Unlock()
}

func (c *cache) HSetNx(key, field string, val any) bool {
	c.hash_mu.RLock()
	hash, ok := c.hashItems[key]
	c.hash_mu.RUnlock()
	if !ok {
		item := HASHItem{
			Object: map[string]any{field: val},
			Key:    key,
		}
		c.hash_mu.Lock()
		c.hashItems[key] = item
		c.hash_mu.Unlock()
		return true
	}
	c.hash_mu.RLock()
	_, ok = hash.Object[field]
	c.hash_mu.RUnlock()
	if ok {
		return false
	}
	c.hash_mu.Lock()
	hash.Object[field] = val
	c.hash_mu.Unlock()
	return true
}

func (c *cache) HDel(key string, fields ...string) {
	if len(fields) == 0 { //全部删除
		c.hash_mu.Lock()
		item, ok := c.hashDelete(key)
		c.hash_mu.Unlock()
		if item.Expiration > 0 && ok {
			c.timeWheel.RemoveTimer(key)
		}
		if ok && item.CallBack && c.deleteCallBack != nil {
			c.deleteCallBack(item.Key, item.Object)
		}
		return
	}
	c.hash_mu.Lock()
	if hash, ok := c.hashItems[key]; ok {
		for _, field := range fields {
			delete(hash.Object, field)
		}
	}
	c.hash_mu.Unlock()
}

func (c *cache) HExists(key string, fields ...string) bool {
	c.hash_mu.RLock()
	hash, ok := c.hashItems[key]
	c.hash_mu.RUnlock()
	if !ok {
		return false
	}
	var res bool = true
	for _, field := range fields {
		if _, ok := hash.Object[field]; !ok {
			res = false
		}
	}
	return res
}

func (c *cache) HGet(key string, fields ...string) map[string]any {
	c.hash_mu.RLock()
	defer c.hash_mu.RUnlock()
	hash, ok := c.hashItems[key]
	res := make(map[string]any, len(hash.Object))
	if !ok {
		return res
	}
	for _, field := range fields {
		if val, ok := hash.Object[field]; ok {
			res[field] = val
		}
	}
	return res
}

func (c *cache) HGetAll(key string) map[string]any {
	c.hash_mu.RLock()
	defer c.hash_mu.RUnlock()
	hash, ok := c.hashItems[key]
	res := make(map[string]any, len(hash.Object))
	if !ok {
		return res
	}
	for field, val := range hash.Object {
		res[field] = val
	}
	return res
}

func (c *cache) HKeys(key string) []string {
	c.hash_mu.RLock()
	defer c.hash_mu.RUnlock()
	hash, ok := c.hashItems[key]
	res := make([]string, 0, len(hash.Object))
	if !ok {
		return res
	}
	for field, _ := range hash.Object {
		res = append(res, field)
	}
	return res
}

func (c *cache) HVAls(key string) []any {
	c.hash_mu.RLock()
	defer c.hash_mu.RUnlock()
	hash, ok := c.hashItems[key]
	res := make([]any, 0, len(hash.Object))
	if !ok {
		return res
	}
	for _, val := range hash.Object {
		res = append(res, val)
	}
	return res
}

func (c *cache) SAdd(key string, d time.Duration, callBack bool, members ...any) {
	if len(members) == 0 {
		return
	}
	var endTime int64
	if d > 0 {
		endTime = time.Now().Add(d).Unix()
	}
	c.set_mu.RLock()
	setItem, ok := c.setItems[key]
	c.set_mu.RUnlock()
	if ok {
		c.set_mu.Lock()
		for _, member := range members {
			if val, ok := setItem.Object[member]; ok {
				if val.Expiration > 0 {
					c.timeWheel.RemoveTimer(val.timeWheelKey)
				}
			}
			timeWheelKey := c.snowflake.Generate().String()
			item := Set{
				Key:          key,
				timeWheelKey: timeWheelKey,
				Member:       member,
				Expiration:   endTime,
				CallBack:     callBack,
			}
			setItem.Object[member] = item
			c.timeWheel.AddTimer(d, timeWheelKey, item)
		}
		c.set_mu.Unlock()
		return
	}
	c.set_mu.Lock()
	object := map[any]Set{}
	for _, member := range members {
		timeWheelKey := c.snowflake.Generate().String()
		item := Set{
			Key:          key,
			timeWheelKey: timeWheelKey,
			Member:       member,
			Expiration:   endTime,
			CallBack:     callBack,
		}
		object[member] = item
		c.timeWheel.AddTimer(d, timeWheelKey, item)
	}
	c.setItems[key] = SetItem{Object: object}
	c.set_mu.Unlock()
}

func (c *cache) setDelete(key string, memberKey any) (Set, bool) {
	if v, ok := c.setItems[key]; ok {
		if set, ok := v.Object[memberKey]; ok {
			delete(v.Object, memberKey)
			return set, true
		} else {
			return Set{}, false
		}
	}
	return Set{}, false
}

func (c *cache) SCard(key string) int {
	c.set_mu.RLock()
	setItem, ok := c.setItems[key]
	c.set_mu.RUnlock()
	if !ok {
		return 0
	}
	return len(setItem.Object)
}

func (c *cache) SRem(key string, members ...any) int {
	i := 0
	if len(members) == 0 {
		return i
	}
	for _, member := range members {
		c.set_mu.Lock()
		item, ok := c.setDelete(key, member)
		c.set_mu.Unlock()
		if item.Expiration > 0 && ok {
			c.timeWheel.RemoveTimer(item.timeWheelKey)
		}
		if ok && item.CallBack && c.deleteCallBack != nil {
			c.deleteCallBack(item.Key, item.Member)
		}
		i++
	}
	return i
}

func (c *cache) SMembers(key string) []any {
	c.set_mu.RLock()
	defer c.set_mu.RUnlock()
	setItem, ok := c.setItems[key]
	members := make([]any, 0, len(setItem.Object))
	if !ok {
		return members
	}
	for member, _ := range setItem.Object {
		members = append(members, member)
	}
	return members
}

func (c *cache) SISMembers(key string, member any) bool {
	c.set_mu.RLock()
	defer c.set_mu.RUnlock()
	setItem, ok := c.setItems[key]
	if !ok {
		return false
	}
	if _, ok := setItem.Object[member]; ok {
		return true
	}
	return false
}
