package login

type LoginResp struct {
	Token           string `json:"token"`           //用户的token
	JumpType        int    `json:"jumpType"`        //跳转类型 1完善信息页面,2选择绑定多账号,3首页
	ChooseUserToken string `json:"chooseUserToken"` //选择用户的token,当jumpType=2的时候该值使用
	Id              string `json:"id"`              //对应的用户id
	Uid32           int32  `json:"uid32"`
}
