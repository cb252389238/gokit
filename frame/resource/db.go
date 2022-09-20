package resource

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbOnce sync.Once
	dbSets *MysqlSets
)

type MysqlSets struct {
	mysql map[string]*gorm.DB
	l     sync.RWMutex
}

func (r *MysqlSets) Key(key ...string) *gorm.DB {
	r.l.RLock()
	defer r.l.RUnlock()
	var name string
	if len(key) <= 0 {
		name = "default"
	} else {
		name = key[0]
	}
	if db, ok := r.mysql[name]; ok {
		conf := GetAllHotConf()
		if conf.Debug {
			return db.Debug()
		} else {
			return db
		}
	}
	return nil
}

func NewDb() *MysqlSets {
	dbOnce.Do(func() {
		conf := GetAllHotConf()
		//载入mysql集合
		dbClients := map[string]*gorm.DB{}
		for _, m := range conf.Mysql {
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Database)
			open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			db, err := open.DB()
			if err != nil {
				panic(err)
			}
			db.SetMaxOpenConns(5)
			db.SetMaxIdleConns(10)
			dbClients[m.Name] = open
		}
		dbSets = &MysqlSets{
			mysql: dbClients,
		}

	})
	return dbSets
}

// GetDb
//
func GetDb(keys ...string) *gorm.DB {
	db := NewDb()
	key := "default"
	if len(keys) > 0 {
		key = keys[0]
	}
	return db.Key(key)
}
