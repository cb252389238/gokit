package user

type UserInfoReq struct {
	UserId string `json:"userId" form:"userId" validate:"required"`
}
