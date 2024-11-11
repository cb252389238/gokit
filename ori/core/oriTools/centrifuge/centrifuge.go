package centrifuge

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var centrifuge *Centrifuge

// 定义离心机结构体
type Centrifuge struct {
	publisher                 *Publisher                                   //订阅发布实例
	platformArray             []string                                     //客户端类型集合
	userNum                   int                                          //用户数量
	connectNum                int                                          //连接数
	cache                     *Cache                                       //本地缓存
	l                         sync.RWMutex                                 //读写锁
	users                     map[string]*UserConnectInfo                  //用户集合
	cacheLifeTime             time.Duration                                //生命周期
	beforeAddCallback         func(userId, platform string) error          //添加前回调
	afterAddCallback          func(userId, platform string) error          //添加后回调
	beforeDelCallback         func(userId, platform string) error          //删除前回调
	afterDelCallback          func(userId, platform string) error          //删除后回调
	beforeSubscribeCallback   func(userId, platform, sinceId string) error //订阅频道前回调
	afterSubscribeCallback    func(userId, platform, sinceId string) error //订阅频道后回调
	beforeUnSubscribeCallback func(userId, platform, sinceId string) error //取消订阅前回调
	afterUnSubscribeCallback  func(userId, platform, sinceId string) error //取消订阅后回调
}

type UserConnectInfo struct {
	userId        string //用户ID
	platform      string //客户端类型
	websocketConn *websocket.Conn
	subscriber    *Subscriber //订阅实例
	sceneId       string      //场景ID
}

// 实例化
func New(platform []string, lifeTime time.Duration) (*Centrifuge, error) {
	if len(platform) == 0 {
		return nil, errors.New("platform is empty")
	}
	if centrifuge != nil {
		return nil, errors.New("centrifuge already exists")
	}
	pubsub, err := NewPublisher(platform)
	if err != nil {
		return nil, err
	}
	c := NewCache(NoExpiration, time.Second*10)
	users := make(map[string]*UserConnectInfo)
	centrifuge = &Centrifuge{
		publisher:     pubsub,
		platformArray: platform,
		cache:         c,
		users:         users,
		cacheLifeTime: lifeTime,
	}
	return centrifuge, nil
}

// 返回实例
func Instance() *Centrifuge {
	return centrifuge
}

// 设置缓存删除过期回调
// k = uid:platform
// v = *UserConnectInfo
func (c *Centrifuge) CacheDeleteCallback(callback func(k string, v any)) {
	c.cache.OnEvicted(callback)
}

// 设置添加用户前回调函数
func (c *Centrifuge) BeforeAddUserCallback(callback func(userId, platform string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.beforeAddCallback = callback
}

// 设置添加用户后回调函数
func (c *Centrifuge) AfterAddUserCallback(callback func(userId, platform string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.afterAddCallback = callback
}

// 删除用户前回调
func (c *Centrifuge) BeforeDelUserCallback(callback func(userId, platform string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.beforeDelCallback = callback
}

// 删除用户后回调
func (c *Centrifuge) AfterDelUserCallback(callback func(userId, platform string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.afterDelCallback = callback
}

// 订阅频道前回调
func (c *Centrifuge) BeforeSubscribeCallback(callback func(userId, platform, sinceId string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.beforeSubscribeCallback = callback
}

// 订阅频道后回调
func (c *Centrifuge) AfterSubscribeCallback(callback func(userId, platform, sinceId string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.afterSubscribeCallback = callback
}

// 取消订阅前回调
func (c *Centrifuge) BeforeUnSubscribeCallback(callback func(userId, platform, sinceId string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.beforeUnSubscribeCallback = callback
}

// 取消订阅后回调
func (c *Centrifuge) AfterUnSubscribeCallback(callback func(userId, platform, sinceId string) error) {
	c.l.Lock()
	c.l.Unlock()
	c.afterUnSubscribeCallback = callback
}

// 用户数量自增
func (c *Centrifuge) incrUserNum(userId string) {
	connectNum := 0
	for _, v := range c.platformArray {
		if _, ok := c.users[getUid(userId, v)]; ok {
			connectNum++
		}
	}
	if connectNum <= 1 {
		c.userNum++
	}
}

// 用户数量自减
func (c *Centrifuge) decrUserNum(userId string) {
	connectNum := 0
	for _, v := range c.platformArray {
		if _, ok := c.users[getUid(userId, v)]; ok {
			connectNum++
		}
	}
	if connectNum == 0 {
		c.userNum--
	}
}

// 获取用户数量
func (c *Centrifuge) GetUserNum() int {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.userNum
}

// 连接数自增
func (c *Centrifuge) incrConnectNum() {
	c.connectNum++
}

// 连接数自减
func (c *Centrifuge) decrConnectNum() {
	c.connectNum--
}

// 获取连接数
func (c *Centrifuge) GetConnectNum() int {
	c.l.RLock()
	defer c.l.RUnlock()
	return c.connectNum
}

// 心跳增加本地缓存生命周期
func (c *Centrifuge) Heartbeat(userId, platform string) {
	c.l.RLock()
	defer c.l.RUnlock()
	uid := getUid(userId, platform)
	if _, ok := c.users[uid]; ok {
		c.cache.Expire(uid, c.cacheLifeTime)
	}
}

// 添加
func (c *Centrifuge) Add(userId, platform string, websocketConn *websocket.Conn) error {
	if userId == "" || platform == "" || websocketConn == nil {
		return errors.New("param error")
	}
	logicFunc := func(userId, platform string, websocketConn *websocket.Conn) error {
		c.l.Lock()
		defer c.l.Unlock()
		uid := getUid(userId, platform)
		subscribe, err := c.publisher.Subscribe(TopicUserOnline, uid, platform)
		if err != nil {
			return err
		}
		user := &UserConnectInfo{
			userId:        userId,
			platform:      platform,
			websocketConn: websocketConn,
			subscriber:    subscribe,
		}
		c.cache.set(getUid(userId, platform), user, c.cacheLifeTime)
		c.users[uid] = user
		c.incrUserNum(userId)
		c.incrConnectNum()
		return nil
	}
	c.l.RLock()
	beforeAddCallback := c.beforeAddCallback
	afterAddCallback := c.afterAddCallback
	c.l.RUnlock()
	if beforeAddCallback != nil {
		err := beforeAddCallback(userId, platform)
		if err != nil {
			return fmt.Errorf("beforeAddCallback err:%+v", err)
		}
	}
	err := logicFunc(userId, platform, websocketConn)
	if err != nil {
		return err
	}
	if afterAddCallback != nil {
		err = c.afterAddCallback(userId, platform)
		if err != nil {
			return fmt.Errorf("afterAddCallback err:%+v", err)
		}
	}
	return nil
}

// 删除用户
func (c *Centrifuge) Del(userId, platform string) error {
	if userId == "" || platform == "" {
		return errors.New("param error")
	}
	logicFunc := func(userId, platform string) error {
		c.l.Lock()
		defer c.l.Unlock()
		uid := getUid(userId, platform)
		c.cache.Delete(uid)
		//退出所有订阅
		subscribe, ok := c.users[uid]
		if ok {
			if subscribe.subscriber != nil {
				for topic, _ := range subscribe.subscriber.Topic {
					c.publisher.UnSubscribe(topic, uid, platform)
				}
			}
		}
		delete(c.users, uid)
		c.decrUserNum(userId)
		c.decrConnectNum()
		if ok {
			if subscribe.websocketConn != nil {
				subscribe.websocketConn.Close()
			}
		}
		return nil
	}
	c.l.RLock()
	beforeDelCallback := c.beforeDelCallback
	afterDelCallback := c.afterDelCallback
	c.l.RUnlock()
	if beforeDelCallback != nil {
		err := beforeDelCallback(userId, platform)
		if err != nil {
			return fmt.Errorf("beforeDelCallback err:%+v", err)
		}
	}
	err := logicFunc(userId, platform)
	if err != nil {
		return err
	}
	if afterDelCallback != nil {
		err = afterDelCallback(userId, platform)
		if err != nil {
			return fmt.Errorf("afterDelCallback err:%+v", err)
		}
	}
	return nil
}

// 用户订阅频道
func (c *Centrifuge) Subscribe(userId, platform, sceneId string) error {
	if userId == "" || platform == "" || sceneId == "" {
		return errors.New("param error")
	}
	logicFunc := func(userId, platform, sceneId string) error {
		c.l.Lock()
		defer c.l.Unlock()
		uid := getUid(userId, platform)
		user, ok := c.users[uid]
		if !ok {
			return errors.New("user not found")
		}
		user.sceneId = sceneId
		_, err := c.publisher.Subscribe(sceneId, uid, platform)
		if err != nil {
			return err
		}
		return nil
	}
	c.l.RLock()
	beforeSubscribeCallback := c.beforeSubscribeCallback
	afterSubscribeCallback := c.afterSubscribeCallback
	c.l.RUnlock()
	if beforeSubscribeCallback != nil {
		err := beforeSubscribeCallback(userId, platform, sceneId)
		if err != nil {
			return fmt.Errorf("beforeSubscribeCallback err:%+v", err)
		}
	}
	err := logicFunc(userId, platform, sceneId)
	if err != nil {
		return err
	}
	if afterSubscribeCallback != nil {
		err := afterSubscribeCallback(userId, platform, sceneId)
		if err != nil {
			return fmt.Errorf("afterSubscribeCallback err:%+v", err)
		}
	}
	return err
}

// 用户取消频道订阅
func (c *Centrifuge) UnSubscribe(userId, platform, sceneId string) error {
	if userId == "" || platform == "" || sceneId == "" {
		return errors.New("param error")
	}
	logicFunc := func(userId, platform, sceneId string) error {
		c.l.Lock()
		defer c.l.Unlock()
		uid := getUid(userId, platform)
		user, ok := c.users[uid]
		if !ok {
			return errors.New("user not found")
		}
		if user.sceneId == sceneId {
			user.sceneId = ""
		}
		return c.publisher.UnSubscribe(sceneId, uid, platform)
	}
	c.l.RLock()
	beforeUnSubscribeCallback := c.beforeUnSubscribeCallback
	afterUnSubscribeCallback := c.afterUnSubscribeCallback
	c.l.RUnlock()
	if beforeUnSubscribeCallback != nil {
		err := beforeUnSubscribeCallback(userId, platform, sceneId)
		if err != nil {
			return fmt.Errorf("beforeUnSubscribeCallback err:%+v", err)
		}
	}
	err := logicFunc(userId, platform, sceneId)
	if err != nil {
		return err
	}
	if afterUnSubscribeCallback != nil {
		err = afterUnSubscribeCallback(userId, platform, sceneId)
		if err != nil {
			return fmt.Errorf("afterUnSubscribeCallback err:%+v", err)
		}
	}
	return nil
}

// 判断用户否订阅频道
func (c *Centrifuge) IsSubscribe(userId, platform, sceneId string) bool {
	uid := getUid(userId, platform)
	c.l.RLock()
	defer c.l.RUnlock()
	user, ok := c.users[uid]
	if !ok {
		return false
	}
	for topic, _ := range user.subscriber.Topic {
		if topic == sceneId {
			return true
		}
	}
	return false
}

// 判断用户是否存在
func (c *Centrifuge) IsUserExist(userId, platform string) bool {
	c.l.RLock()
	defer c.l.RUnlock()
	uid := getUid(userId, platform)
	_, ok := c.users[uid]
	return ok
}

// 向所有用户推送消息
func (c *Centrifuge) PushAll(message any) {
	if message == nil {
		return
	}
	c.publisher.PublishAll(message)
}

// 向指定客户端推送消息
func (c *Centrifuge) PushAllByPlatform(message any, platform ...string) {
	c.publisher.PublishAll(message, platform...)
}

// 向指定用户推送消息
func (c *Centrifuge) PushToUser(userId string, message any) {
	c.publisher.PublishToUser(userId, message)
}

// 向指定的客户端用户推送消息
func (c *Centrifuge) PushToUserByPlatform(userId string, message any, platform ...string) {
	c.publisher.PublishToUser(userId, message, platform...)
}

// 向指定频道推送消息
func (c *Centrifuge) PushToTopic(sceneId string, message any) {
	c.publisher.Publish(sceneId, message)
}

// 向指定频道客户端推送消息
func (c *Centrifuge) PushToTopicByPlatform(sceneId string, message any, platform ...string) {
	c.publisher.Publish(sceneId, message, platform...)
}
