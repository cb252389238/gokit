package request

type UserReq struct {
	Id   int    `json:"id" form:"id" validate:"omitempty,required,min=1"`
	Name string `json:"name" form:"name" validate:"omitempty"`
}
