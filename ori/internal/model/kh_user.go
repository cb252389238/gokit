package model

type KhUser struct {
	ID int `json:"id" gorm:"column:id"` // 用户ID
}

func (m *KhUser) TableName() string {
	return "kh_user"
}
