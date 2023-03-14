# goOri

### go项目基本架子


### 项目结构说明

```
goOri
├─app #应用目录
│  ├─http           #http服务
│  └─ws             #websocket服务 
├─internal          #核心逻辑
│  ├─core           #核心功能 一般不动
│  │  ├─cache       #通用本地缓存
│  │  ├─config      #配置读取
│  │  ├─database    #数据库实例
│  │  ├─log         日志对象
│  │  ├─monitor     #监控模块
│  │  ├─ori         #核心功能初始化
│  │  ├─oriEngine   #核心功能引擎 提供核心功能调用
│  │  ├─oriSignal   #信号模块
│  │  ├─oriTools    #核心工具箱
│  │  ├─pool        #通用连接池
│  │  └─redis       #redis模块
│  ├─dao            #数据访问对象
│  ├─factory        #工厂方法
│  ├─logic          #逻辑对象
│  ├─model          #模型目录
│  └─service        #服务对象
├─logs              #日志目录
├─typedef           #结构体 变量 常量定义目录
└─util              #工具包
    └─tools         #通用工具包

```
### 项目启动
```go
go run main.go -f ./config.yaml

-f 参数可以忽略，默认项目根目录下的config.yaml
也可以自定义配置文件位置，配置名称随意。

可以支持其他配置文件类型，例如：
go run main.go -f ./config.ini
```

### 交叉编译
```go
linux 环境
set GOOS=linux
go build -o main.go 可执行文件名称

windows 环境
set GOOS=windows
go build -o main.go 可执行文件名称.exe
```

### 项目调用说明
```go
1、从入口文件main.go启动
2、
    defer func() {//捕获错误
        if err := recover(); err != nil {
        log.Fatal(err)
    }
    }()
    ori.Start() //启动项目
	
3、进入 internal/core/ori/init.go
读取-f配置参数，载入配置。根据配置需要启动的服务依次启动->配置文件热更新服务->项目资源服务->监听服务->状态报告服务->http->websocket->自定义服务->信号监听服务。
信号监听服务处于阻塞状态，一直到收到信号。
```

### 项目资源服务说明
```go
    项目资源服务初始化了项目需要的主要核心服务，返回一个指针类型，并发安全。可以全局传递使用。
    type OriEngine struct {
        Wg         *sync.WaitGroup //全局同步控制
        Signal     chan os.Signal //全局控制信号
        WsSignal   chan os.Signal //websocket信号
        HttpSignal chan os.Signal //http服务信号
        L          *sync.RWMutex//读写锁
        Context    context.Context//上下文
        Cancel     context.CancelFunc//上下文取消函数
        Mysql      *database.MysqlSets//sql集合
        Redis      *redis.RedisSets//redis集合
        Pool       pool.Pool        //通用连接池
        Factory    *factory.Factory //工厂类
        Log        *log.LocalLogger//日志
        Cache      *cache.Cache//本地缓存
        WebHook    *dingtalk.DingTalk//钉钉通知
    }
    通过 engine := oriEngine.NewOriEngine()  返回一个*OriEngine类型
    
    使用例子：
    oriEngine.Cache.Set("name", "11", time.Second*10)//设置一个本地缓存 10过期
    oriEngine.Log.Debug("%+v",err)//打印debug日志
    oriEngine.Mysql.Key("kaihei").Where("1=1").Find(&user{})//使用kaihei库查询
    oriEngine.Redis.Key("list").Set("name",1,time.Second*10)//设置redis缓存 10秒过期时间
    p, _ := oriEngine.Pool.Get()//通用连接池获取资源
    oriEngine.Pool.Put(p)//释放归还
```

### 其他说明
```go
    使用一些系统资源除了上面说的通过全局传递OriEngine指针还可以直接调用

    cache.New().Set("key", "val", time.Second*10)                                                 //直接通过包调用 设置本地缓存
    config.GetHotConf()                                                                           //获取所有配置
    database.Db("kaihei").Where()                                                                 //调用数据库
    log.LogDebug("%+v", err)                                                                      //打印日志
    monitor.SendNotice("程序挂了")                                                                    //发送钉钉通知
    oriTools.ConcurrencyAdd("/v1/api")                                                            //请求计算，监控系统将进行并发量统计
    oriTools.ConcurrencyDel("/vi/api")                                                            //请求结束记得释放
    oriTools.GetLoaclIp()                                                                         //本地ip
    redis.NewRedis().Key("list").Set()                                                            //使用redis
    s, _ := serC.New([]string{"127.0.0.1"}, "/order/list", "127.0.0.1", 30, context.Background()) //注册服务发现
    s.GetServices()                                                                               //获取可用服务
```
