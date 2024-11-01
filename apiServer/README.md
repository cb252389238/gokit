# yfapi

### api服务
### 项目结构说明

```
├─app
│  ├─handle     handle函数
│  └─middle     中间件
├─config        配置文件
├─core          核心组件
│  ├─coreCache      本地缓存
│  ├─coreConfig     配置加载
│  ├─coreDb         数据库
│  ├─coreJwtToken   token
│  ├─coreLog        日志
│  ├─coreMq         mq组件
│  ├─corePool       通用连接池
│  ├─coreRedis      redis
│  ├─coreSignal     信号管理
│  └─coreTools      核心工具
│      ├─bloomfilter  布隆过滤器
│      ├─cache        基础本地缓存
│      ├─consistentHash 一致性哈希
│      ├─safemap        安全map，解决无限制内存增长问题
│      ├─serC           服务发现 依赖etcd
│      └─snowflake      雪花id生成器
├─internal              核心业务
│  ├─dao                数据曾
│  ├─engine             全局资源加载引擎
│  ├─logic              逻辑层
│  └─model              模型映射
├─logs                  日志目录
├─typedef               结构，常量，变量定义
└─util                  工具包
    └─easy              常用函数封装

```
### 项目启动
```go
go run main.go -f ./config/config.yaml

-f 参数可以忽略，默认项目根目录下的config.yaml
也可以自定义配置文件位置，配置名称随意。


swag 生成文档指令
如果没有swag指令，执行 go install github.com/swaggo/swag/cmd/swag@latest
swag init --exclude .\core\,.\internal\,.\util\ --parseDependency
将docs下方的swagger.json直接导入到apipost中

```
