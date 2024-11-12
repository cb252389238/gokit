package easy

import "testing"

func TestIpFun(t *testing.T) {
	t.Log(GetLocalIp())
	ip := "192.168.7.1"
	toInt := StringIpToInt(ip)
	t.Log(toInt)
	toString := IntIpToString(toInt)
	t.Log(toString)
	stringToInt := Ipv4StringToInt("192.168.7.1")
	t.Log(stringToInt)
	intToString := Ipv4IntToString(stringToInt)
	t.Log(intToString)
}
