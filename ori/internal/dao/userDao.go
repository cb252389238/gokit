package dao

import (
	"ori/core/oriDb"
	"ori/internal/model"
)

type UserDao struct {
}

func (u *UserDao) GetUserInfoById(uid int) (*model.KhUser, error) {
	data := &model.KhUser{}
	err := oriDb.Db().Model(&model.KhUser{}).Where("id = ?", uid).First(data).Error
	return data, err
}
