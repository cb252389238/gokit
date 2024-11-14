package logic

import (
	"ori/internal/dao"
	"ori/oerror"
	"ori/typedef/request/user"
	user2 "ori/typedef/response/user"
)

type User struct {
}

func (u *User) GetUserInfo(req *user.UserInfoReq, resp *user2.UserInfo) error {
	userModle, err := new(dao.UserDao).GetUserInfoById(req.UserId)
	if err != nil {
		return oerror.NewError(oerror.ErrorCodeUserNotExist, nil)
	}
	resp.UserId = userModle.ID
	return nil
}
