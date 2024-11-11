package centrifuge

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var centrifuge *Centrifuge

// 定义离心机结构体
type Centrifuge struct {
	publisher     *Publisher                  //订阅发布实例
	platformArray []string                    //客户端类型集合
	userNum       int                         //用户数量
	connectNum    int                         //连接数
	cache         *Cache                      //本地缓存
	l             sync.RWMutex                //读写锁
	users         map[string]*userConnectInfo //用户集合
}

type userConnectInfo struct {
	userId        string //用户ID
	platform      string //客户端类型
	websocketConn *websocket.Conn
	subscriber    *Subscriber //订阅实例
	sceneId       string      //场景ID
}

// 实例化
func New(platform []string) (*Centrifuge, error) {
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
	users := make(map[string]*userConnectInfo)
	centrifuge = &Centrifuge{
		publisher:     pubsub,
		platformArray: platform,
		cache:         c,
		users:         users,
	}
	return centrifuge, nil
}

// 返回实例
func Instance() *Centrifuge {
	return centrifuge
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

// 添加用户
func (c *Centrifuge) AddUser(userId, platform string, websocketConn *websocket.Conn) error {
	if userId == "" || platform == "" || websocketConn == nil {
		return errors.New("param error")
	}
	uid := getUid(userId, platform)
	subscribe, err := c.publisher.Subscribe(TopicUserOnline, uid, platform)
	if err != nil {
		return err
	}
	user := &userConnectInfo{
		userId:        userId,
		platform:      platform,
		websocketConn: websocketConn,
		subscriber:    subscribe,
	}
	c.l.Lock()
	c.users[uid] = user
	c.incrUserNum(userId)
	c.incrConnectNum()
	c.l.Unlock()
	return nil
}
