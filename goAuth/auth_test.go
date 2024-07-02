package coreAuth

import (
	"fmt"
	"testing"
)

func TestGoAuth_AddRule(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	if err := goAuth.AddRule("enter_room_hidden", "隐身进厅", 1, "开启后隐身进入房间大厅"); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("up_maiwei", "上麦位", 1, "主动上麦位"); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("withdraw", "提现", 2, "提现权限"); err != nil {
		fmt.Println(err)
	}
}

func TestGoAuth_EditRule(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	if err := goAuth.EditRule(1, "enter_room_hidden", "隐身进入房间", 1, ""); err != nil {
		fmt.Println(err)
	}
}

func TestCoreAuth_DeleteRule(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	goAuth.DeleteRule(2)
}

func TestGoAuth_AddRole(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	if err := goAuth.AddRole("主持人", "1,3"); err != nil {
		t.Error(err)
	}
	if err := goAuth.AddRole("麦未嘉宾", "1"); err != nil {
		t.Error(err)
	}
	if err := goAuth.AddRole("歌手", "1"); err != nil {
		t.Error(err)
	}
}

func TestGoAuth_GiveUserRole(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	if err := goAuth.GiveUserRole(1, 1); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(1, 2); err != nil {
		t.Error(err)
	}
}

func TestGoAuth_ShowRoleList(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	list, err := goAuth.ShowRoleList()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestGoAuth_GetRoleRules(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	list, err := goAuth.GetRoleRules(5)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}

func TestCoreAuth_VerifyAuth(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	auth, err := goAuth.VerifyAuth(2, "withdraw")
	fmt.Println(auth, err)
}

func TestCoreAuth_GetUserRules(t *testing.T) {
	goAuth, err := New(CoreAuthConfig{
		UserName: "root",
		PassWord: "root",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "test",
	})
	if err != nil {
		panic(err)
	}
	rules, err := goAuth.GetUserRules(2, 2)
	fmt.Println(rules, err)
}
