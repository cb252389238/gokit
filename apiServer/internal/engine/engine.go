package engine

import (
	"apiServer/core/coreConfig"
	"apiServer/core/coreDb"
	log2 "apiServer/core/coreLog"
	"apiServer/core/corePg"
	pool2 "apiServer/core/corePool"
	"apiServer/core/coreRedis"
	"apiServer/core/coreSnowflake"
	"apiServer/core/coreTools/snowflake"
	"apiServer/util/easy"
	"context"
	"fmt"
	"os"
	"sync"
)

type OriEngine struct {
	Wg         *sync.WaitGroup      //全局阻塞控制
	Signal     chan os.Signal       //全局控制信号
	HttpSignal chan os.Signal       //http信号
	L          *sync.RWMutex        //全局读写锁
	Context    context.Context      //上下文控制
	Cancel     context.CancelFunc   //上下文退出函数
	Db         *coreDb.MysqlSets    //数据库实例
	Redis      *coreRedis.RedisSets //redis实例
	Pool       pool2.Pool           //通用连接池
	Log        *log2.LocalLogger    //日志实例
	Snowflake  *snowflake.Node      //雪花id生成器实例
	Pg         *corePg.PgSqlSets
}

// NewOriEngine
//
//	@Description:	初始化项目资源依赖
//	@return			*OriEngine
func NewOriEngine() *OriEngine {
	cancel, cancelFunc := context.WithCancel(context.Background())
	var redis *coreRedis.RedisSets
	if len(coreConfig.GetHotConf().Redis) >= 1 {
		redis = coreRedis.NewRedis()
	} else {
		redis = nil
	}
	var db *coreDb.MysqlSets
	if len(coreConfig.GetHotConf().Mysql) >= 1 {
		db = coreDb.NewDb()
	} else {
		db = nil
	}
	var pg *corePg.PgSqlSets
	if len(coreConfig.GetHotConf().Pgsql) >= 1 {
		pg = corePg.NewDb()
	} else {
		pg = nil
	}
	ip := easy.GetLocalIp()
	intIp := easy.Ipv4StringToInt(ip)
	node := intIp % 1000
	fmt.Printf("ip:%s,intIp:%d,node:%d\r\n", ip, intIp, node)
	snow, err := coreSnowflake.New(node)
	if err != nil {
		panic(err)
	}
	mysqlConfig := coreConfig.GetHotConf().Mysql
	if len(mysqlConfig) < 1 {
		panic("mysql 配置错误")
	}
	ctx := &OriEngine{
		Wg:         &sync.WaitGroup{},
		Signal:     make(chan os.Signal),
		HttpSignal: make(chan os.Signal),
		L:          &sync.RWMutex{},
		Context:    cancel,
		Cancel:     cancelFunc,
		Db:         db,
		Pg:         pg,
		Redis:      redis,
		Pool: pool2.NewPool(
			func() (interface{}, error) {
				return 1, nil
			},
			func(v interface{}) error {
				return nil
			},
			100,
			100,
			1000,
		),
		Log:       log2.NewLog(),
		Snowflake: snow,
	}
	return ctx
}
