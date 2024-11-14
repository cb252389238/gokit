package model

type User struct {
	ID   string `json:"id" gorm:"column:id"` // 用户ID
	Name string `json:"name" gorm:"column:name"`
}

func (m *User) TableName() string {
	return "user"
}
