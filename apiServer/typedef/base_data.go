package typedef

type HeaderData struct {
	AppVersion  string `json:"appVersion"`  //app版本
	Platform    string `json:"platform"`    //客户端 ios,android,pc
	Models      string `json:"models"`      //登录设备型号
	MachineCode string `json:"machineCode"` //设备号
	Channel     string `json:"channel"`     //渠道
	Ip          string `json:"ip"`          //登录ip
	Address     string `json:"address"`     //登录地址
}

type TokenData struct {
	UserId     string `json:"userId"`     //用户ID
	ClientType string `json:"clientType"` //客户端类型 ios,android,pc
	Mobile     string `json:"mobile"`
}
