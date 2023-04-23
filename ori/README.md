# goOri

### go项目基本架子


### 项目结构说明

```
goOri
├─app #应用目录
│  ├─http           #http服务
│  └─ws             #websocket服务 
├─core           #核心功能 一般不动
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
├─internal          #核心逻辑
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
go run main.go -f ./config.yaml -s 服务1,服务2 -i

-f 参数可以忽略，默认项目根目录下的config.yaml
也可以自定义配置文件位置，配置名称随意。

-s 需要启动得相关常驻服务。

-i 独立启动，加上-i后不再启动http 或者websocket服务


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
	
3、进入 core/ori/init.go
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

    oriCache.New().Set("key", "val", time.Second*10)                                                 //直接通过包调用 设置本地缓存
    oriConfig.GetHotConf()                                                                           //获取所有配置
    oriDb.Db("kaihei").Where()                                                                 //调用数据库
    oriLog.LogDebug("%+v", err)                                                                      //打印日志
    oriMonitor.SendNotice("程序挂了")                                                                    //发送钉钉通知
    oriTools.ConcurrencyAdd("/v1/api")                                                            //请求计算，监控系统将进行并发量统计
    oriTools.ConcurrencyDel("/vi/api")                                                            //请求结束记得释放
    oriTools.GetLoaclIp()                                                                         //本地ip
    oriRedis.NewRedis().Key("list").Set()                                                            //使用redis
    s, _ := serC.New([]string{"127.0.0.1"}, "/order/list", "127.0.0.1", 30, context.Background()) //注册服务发现
    s.GetServices()                                                                               //获取可用服务
```

### php相关函数

```go
项目内封装了一些php常用得函数
位置：core/oriTools/php

调用方法
php.In_array(1, []int{1, 2, 3})

目前支持得函数

array类
Array_merge 多个数组合并为一个
Array_slice 截取数组一部分
Array_diff 返回数组差集
Array_intersect 返回数组交集
Array_key_exists 检查键名是否在数组内
In_array 检查值是否在数组内
Count 数组元素数量
Array_unique 数组去重
Array_rand 数组随机元素
Array_keys 返回数组中得key
ArrayValues 返回数组中得值
Array_flip 反转键值对
Array_reverse 数组反转
Array_count_values 统计元素出现得次数
Shuffle 打乱数组
Array_shift 删除数组中第一个元素
Array_pop 删除最后一个元素
Array_push 尾部插入元素
Array_unshift 头部插入元素

file类
File_exists 文件是否存在
Is_file 是否是文件
Is_dir 是否是目录
Filesize 文件大小
File_put_contents 写入内容
File_get_contents 读取文件内容
Delete 删除文件
Copy 拷贝文件
Rename 文件改名
Mkdir 创建目录
Realpath 绝对路径
Basename 返回路径中文件名
Fclose 关闭文件资源

html类
Html_entity_decode 把html实体转换为字符
Htmlentities html字符转实体

math类
Abs 绝对值
Ceil 向上取整
Floor 向下取整
Max 返回最大值
Min 返回最小值
Mt_rand 返回随机数

net类
Ip2long ip转int
Long2ip int转ip

常用其他函数
Empty 是否为空
Is_numeric 是否数字字符
Exit 退出
Die 退出
Getenv 获取环境变量
Putenv 设置环境变量
Version_compare 版本比较

打印类
Echo
Var_dump

string类
Strlen 返回字符串长度
Mb_strlen 返回包含中文字符长度
Str_replace 字符替换
Explode 分割字符
Implode 合并字符
Substr 返回一部分
Strtolower 转为小写
Strtoupper 转为大写
Strrev 反转字符
Str_repeat 字符重复n次返回新字符串
Str_shuffle 随机打乱字符串
Parse_str 解析url字符串参数
Trim 清楚字符串两边特定字符
Ltrim 清楚左边特定字符
Rtrim 清楚右边特定字符
Json_decode json解析
Json_encode 转json
Sha1 sha1加密
Md5 md5加密
Crc32 crc32算法
Urlencode
Urldecode
Base64_encode
Base64_decode

time类
Strtotime 字符串转时间戳
Date 格式化时间
Checkdate 检查是否是时间格式
Sleep 睡眠
Usleep 睡眠微秒


```
