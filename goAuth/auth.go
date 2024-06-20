package goAuth

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

var coreAuth *CoreAuth

type CoreAuth struct {
	dsn   string
	cache map[string]cacheData
	l     sync.Mutex
	p     Pool
}

type CoreAuthConfig struct {
	UserName string
	PassWord string
	Host     string
	Port     int
	Database string
}

func (c CoreAuthConfig) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", c.UserName, c.PassWord, c.Host, c.Port, c.Database)
}

func Auth() *CoreAuth {
	return coreAuth
}

func New(config CoreAuthConfig) (*CoreAuth, error) {
	if coreAuth != nil {
		return coreAuth, nil
	}
	auth := &CoreAuth{
		dsn:   config.GetDsn(),
		cache: map[string]cacheData{},
	}
	init, err := auth.poolInit()
	if err != nil {
		return nil, errors.New("初始化连接池失败:" + err.Error())
	}
	auth.p = init
	go auth.checkCache()
	fmt.Printf("初始化auth成功   连接池数:%d\r\n", auth.p.Len())
	coreAuth = auth
	return coreAuth, nil
}

func (this *CoreAuth) poolInit() (Pool, error) {
	//factory 创建连接的方法
	factory := func() (any, error) { return sql.Open("mysql", this.dsn) }
	//close 关闭连接的方法
	closed := func(v any) error { return v.(*sql.DB).Close() }
	//ping 检测连接的方法
	ping := func(v any) error { return v.(*sql.DB).Ping() }
	//创建一个连接池： 初始化5，最大空闲连接是20，最大并发连接30
	poolConfig := &Config{
		InitialCap: 10, //资源池初始连接数
		MaxIdle:    10, //最大空闲连接数
		MaxCap:     50, //最大并发连接数
		Factory:    factory,
		Close:      closed,
		Ping:       ping,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	return NewChannelPool(poolConfig)
}

// 添加权限
func (this *CoreAuth) AddRule(name, title string, category int, categoryName string) error {
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

// 修改权限
func (this *CoreAuth) EditRule(id int, name, title string, category int, categoryName string) error {
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
	this.clear()
	return nil
}

// 删除权限
func (this *CoreAuth) DeleteRule(id int) error {
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
	this.clear()
	return nil
}

// 添加新角色
func (this *CoreAuth) AddRole(title string, rules string) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `insert into ` + auth_role_name + ` (title,rules) values(?,?)`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title, rules)
	if err != nil {
		return err
	}
	return nil
}

// 删除角色
func (this *CoreAuth) DeleteRole(id int) error {
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
	this.clear()
	return nil
}

// 修改角色
func (this *CoreAuth) EditRole(id int, title, rules string) error {
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
	this.clear()
	return nil
}

// 分配角色
func (this *CoreAuth) GiveUserRole(userId, roleId int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `insert into ` + auth_role_access_name + ` (user_id,role_id) values(?,?)`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, roleId)
	if err != nil {
		return err
	}
	this.clear()
	return nil
}

// 删除角色
func (this *CoreAuth) DeleteUserRole(userId, groupId int) error {
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `delete from ` + auth_role_access_name + ` where user_id=? and group_id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userId, groupId)
	if err != nil {
		return err
	}
	this.clear()
	return nil
}

// 获取角色列表
func (this *CoreAuth) ShowRoleList() ([]map[string]any, error) {
	key := "cache_role_list"
	c := this.get(key)
	if c != nil {
		return c.([]map[string]any), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select * from ` + auth_role_name + ` where status=1`
	rows, err := db.(*sql.DB).Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	list := []map[string]any{}
	var id int
	var title string
	var status int
	var rules string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &status, &rules); err != nil {
			return nil, err
		}
		list = append(list, map[string]any{"id": id, "title": title, "status": status, "rules": rules})
	}
	this.set(key, list)
	return list, nil
}

// 获取角色下的用户
func (this *CoreAuth) ShowRoleUserList(roleId int) ([]map[string]any, error) {
	key := "cache_role_user_list_" + strconv.Itoa(roleId)
	c := this.get(key)
	if c != nil {
		return c.([]map[string]any), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select user.id as userId from ` + auth_role_name + ` role inner join ` + auth_role_access_name +
		` access on role.id=access.role_id inner join ` + user_table_name + ` user on access.user_id=user.id where role.id=?`
	stmt, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(roleId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	list := []map[string]any{}
	var userId, name any
	for rows.Next() {
		if err := rows.Scan(&userId, &name); err != nil {
			return nil, err
		}
		list = append(list, map[string]any{"userId": userId, "name": name})
	}
	this.set(key, list)
	return list, nil
}

// 获取角色权限
func (this *CoreAuth) GetRoleRules(roleId int) (map[int][]any, error) {
	key := "cache_role_rule_list_" + strconv.Itoa(roleId)
	c := this.get(key)
	if c != nil {
		return c.(map[int][]any), nil
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
	allRules := []map[string]any{}
	var id, category int
	var name, title, categoryName string
	for ruleRows.Next() {
		if err := ruleRows.Scan(&id, &name, &title, &category, &categoryName); err != nil {
			return nil, err
		}
		allRules = append(allRules, map[string]any{"id": id, "name": name, "title": title, "category": category, "categoryName": categoryName, "select": 0})
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
	var sortList = map[int][]any{}
	for _, v := range allRules {
		if category, ok := v["category"].(int); ok {
			sortList[category] = append(sortList[category], v)
		}
	}
	this.set(key, sortList)
	return sortList, nil
}

// 验证用户是否拥有权限
func (this *CoreAuth) VerifyAuth(userId int, ruleName string) (bool, error) {
	rules, err := this.getUserRules(userId, 0)
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

// 获取用户的所有权限
// category 0全部 1前端权限 2后端权限
func (this *CoreAuth) getUserRules(userId int, category int) ([]string, error) {
	key := "cache_user_rule_list_" + strconv.Itoa(userId)
	c := this.get(key)
	if c != nil {
		return c.([]string), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select role.rules from ` + auth_role_access_name + ` access inner join ` + auth_role_name + ` role on access.role_id=role.id where role.status=1 and access.user_id=?`
	stms, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stms.Query(userId)
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
	var ruleRows *sql.Rows
	if category == 0 {
		query = `select name from ` + auth_rule_name + ` where id in(` + strings.Join(ruleSlice, ",") + `)`
		ruleRows, err = db.(*sql.DB).Query(query)
	} else {
		query = `select name from ` + auth_rule_name + ` where id in(` + strings.Join(ruleSlice, ",") + `) and category=?`
		stmt, err := db.(*sql.DB).Prepare(query)
		if err != nil {
			return nil, err
		}
		ruleRows, err = stmt.Query(category)
	}
	if err != nil {
		return nil, err
	}
	defer ruleRows.Close()
	list := []string{}
	for ruleRows.Next() {
		ruleRows.Scan(&ruleStr)
		list = append(list, ruleStr)
	}
	this.set(key, list)
	return list, err
}

func (this *CoreAuth) GetUserRules(userId int, category int) (map[string]any, error) {
	key := "cache_user_rule_list_map_" + strconv.Itoa(userId)
	c := this.get(key)
	if c != nil {
		return c.(map[string]any), nil
	}
	db, err := this.p.Get()
	defer this.p.Put(db)
	query := `select role.rules from ` + auth_role_access_name + ` access inner join ` + auth_role_name + ` role on access.role_id=role.id where role.status=1 and access.user_id=?`
	stms, err := db.(*sql.DB).Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stms.Query(userId)
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
	ruleIdsMap := map[string]struct{}{}
	for _, v := range ruleSlice {
		ruleIdsMap[v] = struct{}{}
	}
	var ruleRows *sql.Rows
	if category == 0 {
		query = `select id,name from ` + auth_rule_name + ` where status=1`
		ruleRows, err = db.(*sql.DB).Query(query)
	} else {
		query = `select id,name from ` + auth_rule_name + ` where status=1 and category=?`
		stmt, err := db.(*sql.DB).Prepare(query)
		if err != nil {
			return nil, err
		}
		ruleRows, err = stmt.Query(category)
	}
	if err != nil {
		return nil, err
	}
	defer ruleRows.Close()
	var name string
	var id int
	list := map[string]any{}
	for ruleRows.Next() {
		ruleRows.Scan(&id, &name)
		if _, ok := ruleIdsMap[strconv.Itoa(id)]; ok {
			list[name] = 1
		} else {
			list[name] = 0
		}
	}
	this.set(key, list)
	return list, err
}

type cacheData struct {
	expires int64
	data    interface{}
}

func (this *CoreAuth) get(key string) any {
	this.l.Lock()
	defer this.l.Unlock()
	data := this.cache[key]
	if time.Now().Unix() > data.expires {
		delete(this.cache, key)
		return nil
	}
	return data.data
}

func (this *CoreAuth) set(key string, value any) {
	this.l.Lock()
	defer this.l.Unlock()
	data := cacheData{
		expires: time.Now().Unix() + 120,
		data:    value,
	}
	this.cache[key] = data
}

func (this *CoreAuth) del(key string) {
	this.l.Lock()
	defer this.l.Unlock()
	delete(this.cache, key)
}

func (this *CoreAuth) checkCache() {
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

func (this *CoreAuth) clear() {
	this.l.Lock()
	defer this.l.Unlock()
	this.cache = make(map[string]cacheData)
}
