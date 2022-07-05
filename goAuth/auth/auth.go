package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type GoAuth struct {
	dsn   string
	cache map[string]cacheData
	l     sync.Mutex
	p     Pool
}

func New(dsn string) (*GoAuth, error) {
	goauth := &GoAuth{
		dsn:   dsn,
		cache: map[string]cacheData{},
	}
	init, err := goauth.poolInit()
	if err != nil {
		return nil, errors.New("初始化连接池失败:" + err.Error())
	}
	goauth.p = init
	go goauth.checkCache()
	fmt.Printf("初始化auth成功   连接池数:%d\r\n", goauth.p.Len())
	return goauth, nil
}

func (this *GoAuth) poolInit() (Pool, error) {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return sql.Open("mysql", this.dsn) }
	//close 关闭连接的方法
	closed := func(v interface{}) error { return v.(*sql.DB).Close() }
	//ping 检测连接的方法
	ping := func(v interface{}) error { return v.(*sql.DB).Ping() }
	//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
	poolConfig := &Config{
		InitialCap: 5,  //资源池初始连接数
		MaxIdle:    5,  //最大空闲连接数
		MaxCap:     10, //最大并发连接数
		Factory:    factory,
		Close:      closed,
		Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	return NewChannelPool(poolConfig)
}

/**
添加权限规则
*/
func (this *GoAuth) AddRule(name, title string, category int, categoryName string) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `insert into ` + auth_rule_name + ` (name,title,category,categoryName) values(?,?,?,?)`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, title, category, categoryName)
	if err != nil {
		return err
	}
	return nil
}

/**
修改权限规则
*/
func (this *GoAuth) EditRule(id int, name, title string, category int, categoryName string) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `update ` + auth_rule_name + ` set name=?,title=?,category=?,categoryName=? where id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, title, category, categoryName, id)
	if err != nil {
		return err
	}
	return nil
}

/**
删除权限规则
*/
func (this *GoAuth) DeleteRule(id int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `delete from ` + auth_rule_name + ` where id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

/**
添加新角色
*/
func (this *GoAuth) AddRole(title string) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `insert into ` + auth_role_name + ` (title) values(?)`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

/**
删除角色
*/
func (this *GoAuth) DeleteRole(id int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `delete from ` + auth_role_name + ` where id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

/**
修改角色
*/
func (this *GoAuth) EditRole(id int, title, rules string) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `update ` + auth_role_name + ` set title=?,rules=? where id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, rules, id)
	if err != nil {
		return err
	}
	return nil
}

/**
给用户赋予角色
*/
func (this *GoAuth) GiveUserRole(uid, roleId int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `insert into ` + auth_role_access_name + ` (uid,role_id) values(?,?)`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid, roleId)
	if err != nil {
		return err
	}
	return nil
}

/**
删除用户角色
*/
func (this *GoAuth) DeleteUserRole(uid, groupId int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `delete from ` + auth_role_access_name + ` where uid=? and group_id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid, groupId)
	if err != nil {
		return err
	}
	return nil
}

/**
获取角色列表
*/
func (this *GoAuth) ShowRoleList() ([]map[string]interface{}, error) {
	key := "cache_role_list"
	c := this.get(key)
	if c != nil {
		return c.([]map[string]interface{}), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select * from ` + auth_role_name + ` where status=1`
	rows, err := db.(*sql.DB).Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	list := []map[string]interface{}{}
	var id, title, status, rules interface{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &status, &rules); err != nil {
			return nil, err
		}
		list = append(list, map[string]interface{}{"id": id, "title": title, "status": status, "rules": rules})
	}
	this.set(key, list)
	return list, nil
}

/**
获取角色下的用户
*/
func (this *GoAuth) ShowRoleUserList(roleId int) ([]map[string]interface{}, error) {
	key := "cache_role_user_list_" + strconv.Itoa(roleId)
	c := this.get(key)
	if c != nil {
		return c.([]map[string]interface{}), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select user.id as uid,user.name from ` + auth_role_name + ` role inner join ` + auth_role_access_name +
		` access on role.id=access.role_id inner join ` + user_table_name + ` user on access.uid=user.id where role.id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(roleId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	list := []map[string]interface{}{}
	var uid, name interface{}
	for rows.Next() {
		if err := rows.Scan(&uid, &name); err != nil {
			return nil, err
		}
		list = append(list, map[string]interface{}{"uid": uid, "name": name})
	}
	this.set(key, list)
	return list, nil
}

/**
获取角色权限
*/
func (this *GoAuth) GetRoleRules(roleId int) (map[int][]interface{}, error) {
	key := "cache_role_rule_list_" + strconv.Itoa(roleId)
	c := this.get(key)
	if c != nil {
		return c.(map[int][]interface{}), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select rules from ` + auth_role_name + ` where status=1 and id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	roleRows := stmt.QueryRow(roleId)
	var rules string
	err = roleRows.Scan(&rules)
	if err != nil {
		return nil, err
	}
	ruleIds := strings.Split(rules, ",")
	query = `select id,name,title,category,categoryName from ` + auth_rule_name
	ruleRows, err := db.(*sql.DB).Query(query)
	defer ruleRows.Close()
	if err != nil {
		return nil, err
	}
	allRules := []map[string]interface{}{}
	var id, category int
	var name, title, categoryName string
	for ruleRows.Next() {
		if err := ruleRows.Scan(&id, &name, &title, &category, &categoryName); err != nil {
			return nil, err
		}
		allRules = append(allRules, map[string]interface{}{"id": id, "name": name, "title": title, "category": category, "categoryName": categoryName, "select": 0})
	}
	for k, rule := range allRules {
		if id, ok := rule["id"].(int); ok {
			for _, ruleId := range ruleIds {
				idid, _ := strconv.Atoi(ruleId)
				if idid == id {
					allRules[k]["select"] = 1
				}
			}
		}
	}
	var sortList = map[int][]interface{}{}
	for _, v := range allRules {
		if category, ok := v["category"].(int); ok {
			sortList[category] = append(sortList[category], v)
		}
	}
	this.set(key, sortList)
	return sortList, nil
}

/**
验证用户是否拥有权限
*/
func (this *GoAuth) VerifyAuth(uid int, ruleName string) (bool, error) {
	rules, err := this.getUserRules(uid)
	var rback bool = false
	if err != nil {
		return false, err
	}
	for _, v := range rules {
		if v == ruleName {
			rback = true
			break
		}
	}
	return rback, nil
}

func (this *GoAuth) getUserRules(uid int) ([]string, error) {
	key := "cache_user_rule_list_" + strconv.Itoa(uid)
	c := this.get(key)
	if c != nil {
		return c.([]string), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select role.rules from ` + auth_role_access_name + ` access inner join ` + auth_role_name + ` role on access.role_id=role.id where role.status=1 and access.uid=?`
	stms, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stms.Query(uid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var ruleStr string
	ruleSlice := []string{}
	for rows.Next() {
		if err := rows.Scan(&ruleStr); err != nil {
			return nil, err
		}
		ruleSlice = append(ruleSlice, strings.Split(ruleStr, ",")...)
	}
	query = `select name from ` + auth_rule_name + ` where id in(` + strings.Join(ruleSlice, ",") + `)`
	ruleRows, err := db.(*sql.DB).Query(query)
	defer ruleRows.Close()
	if err != nil {
		return nil, err
	}
	list := []string{}
	for ruleRows.Next() {
		ruleRows.Scan(&ruleStr)
		list = append(list, ruleStr)
	}
	this.set(key, list)
	return list, err
}

type cacheData struct {
	expires int64
	data    interface{}
}

func (this *GoAuth) get(key string) interface{} {
	this.l.Lock()
	defer this.l.Unlock()
	data := this.cache[key]
	if time.Now().Unix() > data.expires {
		delete(this.cache, key)
		return nil
	}
	return data.data
}

func (this *GoAuth) set(key string, value interface{}) {
	this.l.Lock()
	defer this.l.Unlock()
	data := cacheData{
		expires: time.Now().Unix() + 120,
		data:    value,
	}
	this.cache[key] = data
}

func (this *GoAuth) del(key string) {
	this.l.Lock()
	defer this.l.Unlock()
	delete(this.cache, key)
}

func (this *GoAuth) checkCache() {
	for {
		this.l.Lock()
		for key, value := range this.cache {
			if time.Now().Unix() > value.expires {
				delete(this.cache, key)
			}
		}
		this.l.Unlock()
		time.Sleep(time.Second * 300)
	}
}
