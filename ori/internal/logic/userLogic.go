package logic

import (
	"ori/core/oriLog"
	"ori/internal/dao"
	"ori/internal/model"
)

type User struct {
	Uid int
}

func (u *User) GetUserInfo() *model.KhUser {
	user, err := new(dao.UserDao).GetUserInfoById(u.Uid)
	if err != nil {
		oriLog.Error("%+v", err)
		return nil
	}
	return user
}
