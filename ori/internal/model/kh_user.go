package model

type KhUser struct {
	ID                  int    `json:"id" gorm:"column:id"`                                         // 用户ID
	Pid                 int    `json:"pid" gorm:"column:pid"`                                       // 父级ID
	InvateCode          string `json:"invate_code" gorm:"column:invate_code"`                       // 邀请码
	OldID               int    `json:"old_id" gorm:"column:old_id"`                                 // 合并后ID：0-未合并，大于0-合并后ID
	Username            string `json:"username" gorm:"column:username"`                             // 用户名
	Password            string `json:"password" gorm:"column:password"`                             // 密码
	Email               string `json:"email" gorm:"column:email"`                                   // 用户邮箱
	PhoneCode           uint32 `json:"phone_code" gorm:"column:phone_code"`                         // 手机号国际区号，默认中国大陆86
	Mobile              string `json:"mobile" gorm:"column:mobile"`                                 // 用户手机
	CreateTime          uint   `json:"create_time" gorm:"column:create_time"`                       // 注册时间
	RegIp               string `json:"reg_ip" gorm:"column:reg_ip"`                                 // 注册IP
	LastLoginTime       uint   `json:"last_login_time" gorm:"column:last_login_time"`               // 最后登录时间
	LastLoginIp         string `json:"last_login_ip" gorm:"column:last_login_ip"`                   // 最后登录IP
	UpdateTime          uint   `json:"update_time" gorm:"column:update_time"`                       // 更新时间
	OfflineTime         int    `json:"offline_time" gorm:"column:offline_time"`                     // 离线时间
	Status              int8   `json:"status" gorm:"column:status"`                                 // 用户状态：0-禁用，1-启用，2-合并，3-注销
	GroupID             int16  `json:"group_id" gorm:"column:group_id"`                             // 用户组ID
	Score               int32  `json:"score" gorm:"column:score"`                                   // 用户积分
	Newpm               int16  `json:"newpm" gorm:"column:newpm"`                                   // 新短消息数量
	FreezeReason        string `json:"freeze_reason" gorm:"column:freeze_reason"`                   // 冻结原因
	WhoOperate          int    `json:"who_operate" gorm:"column:who_operate"`                       // 操作人UID
	DeviceID            string `json:"device_id" gorm:"column:device_id"`                           // 0
	Role                uint   `json:"role" gorm:"column:role"`                                     // 角色ID
	VipUid              int    `json:"vip_uid" gorm:"column:vip_uid"`                               // 靓号id
	Remark              string `json:"remark" gorm:"column:remark"`                                 // 账号禁用备注
	Source              string `json:"source" gorm:"column:source"`                                 // 注册来源，默认电竞
	VoidAccountMobile   string `json:"void_account_mobile" gorm:"column:void_account_mobile"`       // 注销手机号
	MobileStatus        int8   `json:"mobile_status" gorm:"column:mobile_status"`                   // 手机号状态 1正常 2异常
	LastLoginProvinceID uint   `json:"last_login_province_id" gorm:"column:last_login_province_id"` // 最后一次登录的省id
	LastLoginCityID     uint   `json:"last_login_city_id" gorm:"column:last_login_city_id"`         // 最后一次登录的市id
	Brand               string `json:"brand" gorm:"column:brand"`                                   // 品牌,1：官网，默认值，多品牌时后面,追加
}

func (m *KhUser) TableName() string {
	return "kh_user"
}
