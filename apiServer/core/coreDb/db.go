package coreDb

import (
	"apiServer/core/coreConfig"
	"apiServer/core/coreLog"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	sets *MysqlSets
)

type MysqlSets struct {
	mysql map[string]*gorm.DB
	l     sync.RWMutex
}

func (r *MysqlSets) Db(key ...string) *gorm.DB {
	r.l.RLock()
	defer r.l.RUnlock()
	var name string
	if len(key) <= 0 {
		name = "default"
	} else {
		name = key[0]
	}
	if db, ok := r.mysql[name]; ok {
		conf := coreConfig.GetHotConf()
		if conf.Debug {
			return db.Debug()
		} else {
			return db
		}
	}
	return nil
}

func NewDb() *MysqlSets {
	once.Do(func() {
		conf := coreConfig.GetHotConf()
		//载入mysql集合
		dbClients := map[string]*gorm.DB{}
		logrusLogger := &LogrusLogger{log: coreLog.NewLog()}
		for _, m := range conf.Mysql {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Database)
			open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: logrusLogger,
			})
			if err != nil {
				panic("数据库初始化失败:" + err.Error())
			}
			db, err := open.DB()
			if err != nil {
				panic(err)
			}
			db.SetMaxOpenConns(5)
			db.SetMaxIdleConns(10)
			dbClients[m.Name] = open
		}
		sets = &MysqlSets{
			mysql: dbClients,
		}

	})
	return sets
}

func db(keys ...string) *gorm.DB {
	db := NewDb()
	key := "default"
	if len(keys) > 0 {
		key = keys[0]
	}
	return db.Db(key)
}

// GetMasterDb
//
//	@Description:	获取主库
//	@return			*gorm.DB
func GetMasterDb() *gorm.DB {
	return db("master")
}

// GetSlaveDb
//
//	@Description:	获取从库
//	@return			*gorm.DB
func GetSlaveDb() *gorm.DB {
	return db("slave")
}
