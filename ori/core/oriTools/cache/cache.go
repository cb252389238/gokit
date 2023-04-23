package cache

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	// 永不过期
	NoExpiration time.Duration = -1
	//默认设置得过期时间
	DefaultExpiration time.Duration = 0
)

// 具体缓存数据结构体
type Item struct {
	Object     any   //数据
	Expiration int64 //过期时间
}

// 判断是否过期
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type Cache struct {
	*cache
}

// 缓存结构体
type cache struct {
	defaultExpiration time.Duration     //默认过期结时间
	items             map[string]Item   //存储缓存数据
	mu                sync.RWMutex      //读写锁
	onEvicted         func(string, any) //回调函数
	janitor           *janitor
}

// 返回新得缓存实例
// defaultExpiration 默认过期时间 小于
// cleanupInterval 清理过期缓存得间隔时间
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	return newCacheWithJanitor(defaultExpiration, cleanupInterval, items)
}

// 创建一个缓存实例内部使用 外部不可达
func newCache(de time.Duration, m map[string]Item) *cache {
	if de == 0 {
		de = -1
	}
	c := &cache{
		defaultExpiration: de,
		items:             m,
	}
	return c
}

// 创建一个缓存伴随一个定时清理任务
func newCacheWithJanitor(de time.Duration, ci time.Duration, m map[string]Item) *Cache {
	c := newCache(de, m)
	C := &Cache{c}
	if ci > 0 {
		runJanitor(c, ci) //执行清理工  定时清理过期缓存
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

func (j *janitor) Run(c *cache) {
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

// 删除过期项目
func (c *cache) DeleteExpired() {
	var evictedItems []keyAndValue
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			ov, evicted := c.delete(k)
			if evicted {
				evictedItems = append(evictedItems, keyAndValue{k, ov})
			}
		}
	}
	c.mu.Unlock()
	for _, v := range evictedItems {
		c.onEvicted(v.key, v.value)
	}
}

// 停止定时清理任务
func stopJanitor(c *Cache) {
	c.janitor.stop <- true
}

func runJanitor(c *cache, ci time.Duration) {
	j := &janitor{
		Interval: ci,
		stop:     make(chan bool),
	}
	c.janitor = j
	go j.Run(c)
}

// 设置缓存
func (c *cache) Set(k string, x any, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[k] = Item{
		Object:     x,
		Expiration: e,
	}
	c.mu.Unlock()
}

// 内部使用不加锁
func (c *cache) set(k string, x any, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Object:     x,
		Expiration: e,
	}
}

// 更新过期时间
func (c *cache) Expire(k string, d time.Duration) error {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return fmt.Errorf("Item %s already exists", k)
	}
	c.mu.RUnlock()
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	e = time.Now().Add(d).UnixNano()
	item.Expiration = e
	c.mu.Lock()
	c.items[k] = item
	c.mu.Unlock()
	return nil
}

// 添加一个值到缓存 过期或者不存在才成功，否则报错
func (c *cache) SetNx(k string, x any, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s already exists", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	return nil
}

// 获取值
func (c *cache) Get(k string) (any, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, false
		}
	}
	c.mu.RUnlock()
	return item.Object, true
}

// 以秒为单位返回过期时间
func (c *cache) Ttl(k string) int64 {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return 0
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return 0
		}
		c.mu.RUnlock()
		return time.Unix(0, item.Expiration).Unix()
	}
	c.mu.RUnlock()
	return 0
}

// 以毫秒为单位返回过期时间
func (c *cache) Pttl(k string) int64 {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return 0
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return 0
		}
		c.mu.RUnlock()
		return time.Unix(0, item.Expiration).UnixMilli()
	}
	c.mu.RUnlock()
	return 0
}

// 内部获取缓存 不加锁
func (c *cache) get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Object, true
}

// 对缓存进行加操作
func (c *cache) Increment(k string, n any) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("Item %s not found", k)
	}
	switch v.Object.(type) {
	case int:
		if s, ok := n.(int); ok {
			v.Object = v.Object.(int) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}

	case int8:
		if s, ok := n.(int8); ok {
			v.Object = v.Object.(int8) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int16:
		if s, ok := n.(int16); ok {
			v.Object = v.Object.(int16) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int32:
		if s, ok := n.(int32); ok {
			v.Object = v.Object.(int32) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int64:
		if s, ok := n.(int64); ok {
			v.Object = v.Object.(int64) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint:
		if s, ok := n.(uint); ok {
			v.Object = v.Object.(uint) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uintptr:
		if s, ok := n.(uintptr); ok {
			v.Object = v.Object.(uintptr) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint8:
		if s, ok := n.(uint8); ok {
			v.Object = v.Object.(uint8) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint16:
		if s, ok := n.(uint16); ok {
			v.Object = v.Object.(uint16) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint32:
		if s, ok := n.(uint32); ok {
			v.Object = v.Object.(uint32) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint64:
		if s, ok := n.(uint64); ok {
			v.Object = v.Object.(uint64) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case float32:
		if s, ok := n.(float32); ok {
			v.Object = v.Object.(float32) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case float64:
		if s, ok := n.(float64); ok {
			v.Object = v.Object.(float64) + s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	default:
		c.mu.Unlock()
		return fmt.Errorf("The value for %s is not an integer", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// 对缓存进行减操作
func (c *cache) Decrement(k string, n any) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("Item not found")
	}
	switch v.Object.(type) {
	case int:
		if s, ok := n.(int); ok {
			v.Object = v.Object.(int) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int8:
		if s, ok := n.(int8); ok {
			v.Object = v.Object.(int8) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int16:
		if s, ok := n.(int16); ok {
			v.Object = v.Object.(int16) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int32:
		if s, ok := n.(int32); ok {
			v.Object = v.Object.(int32) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case int64:
		if s, ok := n.(int64); ok {
			v.Object = v.Object.(int64) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint:
		if s, ok := n.(uint); ok {
			v.Object = v.Object.(uint) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uintptr:
		if s, ok := n.(uintptr); ok {
			v.Object = v.Object.(uintptr) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint8:
		if s, ok := n.(uint8); ok {
			v.Object = v.Object.(uint8) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint16:
		if s, ok := n.(uint16); ok {
			v.Object = v.Object.(uint16) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint32:
		if s, ok := n.(uint32); ok {
			v.Object = v.Object.(uint32) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case uint64:
		if s, ok := n.(uint64); ok {
			v.Object = v.Object.(uint64) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case float32:
		if s, ok := n.(float32); ok {
			v.Object = v.Object.(float32) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	case float64:
		if s, ok := n.(float64); ok {
			v.Object = v.Object.(float64) - s
		} else {
			c.mu.Unlock()
			return fmt.Errorf("The incr value and cache type are not accompanied")
		}
	default:
		c.mu.Unlock()
		return fmt.Errorf("The value for %s is not an integer", k)
	}
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// 删除缓存
func (c *cache) Delete(k string) {
	c.mu.Lock()
	v, evicted := c.delete(k)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(k, v)
	}
}

// 内部使用不加锁
func (c *cache) delete(k string) (any, bool) {
	if c.onEvicted != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v.Object, true
		}
	}
	delete(c.items, k)
	return nil, false
}

type keyAndValue struct {
	key   string
	value any
}

// 设置回调函数，删除或者过期时触发
func (c *cache) OnEvicted(f func(string, any)) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

// 将所有缓存保存到io
func (c *cache) Save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		gob.Register(v.Object)
	}
	err = enc.Encode(&c.items)
	return
}

// 将缓存写入到指定文件
func (c *cache) SaveFile(fname string) error {
	fp, err := os.Create(fname)
	if err != nil {
		return err
	}
	err = c.Save(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

// 载入缓存
func (c *cache) Load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := map[string]Item{}
	err := dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range items {
			ov, found := c.items[k]
			if !found || ov.Expired() {
				c.items[k] = v
			}
		}
	}
	return err
}

// 通过指定文件载入缓存
func (c *cache) LoadFile(fname string) error {
	fp, err := os.Open(fname)
	if err != nil {
		return err
	}
	err = c.Load(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

// 返回所有未过期得缓存
func (c *cache) Items() map[string]Item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]Item, len(c.items))
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if v.Expiration > 0 {
			if now > v.Expiration {
				continue
			}
		}
		m[k] = v
	}
	return m
}

// 返回所有得缓存数量
func (c *cache) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}

// 删除所有得缓存
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = map[string]Item{}
	c.mu.Unlock()
}
