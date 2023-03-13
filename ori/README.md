# ori

### go项目基本架子


### 项目结构说明

```
ori
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