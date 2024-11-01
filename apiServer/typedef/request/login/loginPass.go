package login

type LoginPassReq struct {
	Mobile     string `json:"mobile" validate:"required"`     //手机号
	Password   string `json:"password" validate:"required"`   //密码
	RegionCode string `json:"regionCode" validate:"required"` //区号
}
