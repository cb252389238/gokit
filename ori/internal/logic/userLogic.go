package logic

import (
	"errors"
	"ori/core/oriLog"
	"ori/internal/dao"
	"ori/typedef/request/user"
	user2 "ori/typedef/response/user"
)

type User struct {
}

func (u *User) GetUserInfo(req *user.UserInfoReq, resp *user2.UserInfo) error {
	return errors.New("custom err")
	//return oerror.NewError(oerror.ErrorCodeOperationFrequent, map[string]any{"num": 1})
	//panic("panic test")
	//panic(oerror.ErrorCodeParam)
	err := new(dao.UserDao).GetUserInfoById(req.UserId)
	if err != nil {
		oriLog.Error("%+v", err)
		return err
	}
	resp.UserId = req.UserId
	return nil
}
