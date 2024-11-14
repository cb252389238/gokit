package dao

import (
	"ori/core/oriDb"
	"ori/internal/model"
)

type UserDao struct {
}

func (u *UserDao) GetUserInfoById(uid string) (*model.User, error) {
	res := &model.User{}
	err := oriDb.New().Db("default").Where("id", uid).First(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
