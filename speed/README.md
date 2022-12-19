## speed
#### 快速缓存库，支持超时设置，超时/删除回调


## 安装

```
go get github.com/cb252389238/speed
```




```go
c, err := New()
if err != nil {
    panic(err)
}
//绑定回调删除，当元素过期、被删除得时候触发。v是对应得缓存值
c.BindDeleteCallBackFunc(func(v interface{}) {
    fmt.Println("触发回调函数", v)
})
//设置普通缓存
//k:键 v:值 d:过期时间 callBack:是否触发回调
c.Set(k string, v interface{}, d time.Duration, callBack bool)
//当key不存在时设置成功 否则失败
//k:键 v:值 d:过期时间 callBack:是否触发回调
//false 失败 true 成功
c.SetNx(k string, v interface{}, d time.Duration, callBack bool)bool
//获取缓存值
//k:键
c.GetGet(k string) (interface{}, bool)
//获取值和过期时间
c.GetEx(k string) (interface{}, time.Time, bool)
//删除缓存
c.Del(k string)
//获取所有普通缓存值
c.Items() map[string]interface{}
//获取缓存数量
c.ItemCount() int 
//判断缓存是否存在
c.Exists(k string) bool

//hash设置
c.HSet(key, field string, val interface{})
//为hash设置过期时间
c.HSetEx(key string, d time.Duration, callBack bool)bool
//设置多个字段值
c.HMSet(key string, data map[string]interface{})
//hash字段不存在设置成功，否则失败
c.HSetNx(key, field string, val interface{}) bool
//删除hash  fields为空删除整个hash  否则删除对应得字段
c.HDel(key string, fields ...string)
//判断hash是否存在字段，存在为true。多个字段全部都存在为true 否则为false
c.HExists(key string, fields ...string) bool
//获取hash字段值
c.HGet(key string, fields ...string) map[string]interface{}
//获取hash所有字段值
c.HGetAll(key string) map[string]interface{}
//获取hash所有field值
c.HKeys(key string) []string
//获取hash所有字段值
c.HVAls(key string) []interface{}


//无序集合添加值
c.SAdd(key string, d time.Duration, callBack bool, members ...interface{})
//获取无序集合成员个数
c.SCard(key string) int
//删除无序集合成员  返回删除个数
c.SRem(key string, members ...interface{}) int
//获取所有无序集合成员
c.SMembers(key string) []interface{}
//判断成员是否包含在无序集合中
c.SISMembers(key string, member interface{}) bool
```
