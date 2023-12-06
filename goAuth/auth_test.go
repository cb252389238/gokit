package goAuth

import (
	"fmt"
	"testing"
)

func TestGoAuth_AddRule(t *testing.T) {
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err := goAuth.AddRule("/admin/index", "首页", 0, ""); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("/admin/list", "列表", 0, ""); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("/admin/add", "添加", 0, ""); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("/admin/edit", "修改", 0, ""); err != nil {
		fmt.Println(err)
	}
	if err := goAuth.AddRule("/admin/delete", "删除", 0, ""); err != nil {
		fmt.Println(err)
	}
}

func TestGoAuth_EditRule(t *testing.T) {
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err := goAuth.EditRule(1, "/admin/index", "首页1", 0, ""); err != nil {
		fmt.Println(err)
	}
}

func TestGoAuth_AddRole(t *testing.T) {
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err := goAuth.AddRole("管理员"); err != nil {
		t.Error(err)
	}
	if err := goAuth.AddRole("运营"); err != nil {
		t.Error(err)
	}
	if err := goAuth.AddRole("测试"); err != nil {
		t.Error(err)
	}
}

func TestGoAuth_GiveUserRole(t *testing.T) {
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	if err := goAuth.GiveUserRole(1, 1); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(1, 2); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(1, 3); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(1, 4); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(1, 5); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(2, 1); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(2, 2); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(2, 3); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(3, 1); err != nil {
		t.Error(err)
	}
	if err := goAuth.GiveUserRole(3, 2); err != nil {
		t.Error(err)
	}
}

func TestGoAuth_ShowRoleList(t *testing.T) {
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
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
	goAuth, err := New("root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	list, err := goAuth.GetRoleRules(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(list)
}
