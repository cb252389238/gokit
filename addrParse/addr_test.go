package addr

import (
	"fmt"
	"strings"
	"testing"
)

var blackAddress = []string{"福建-泉州", "福建-漳州", "厦门-思明", "新疆", "西藏", "吉林", "辽宁"}

type addrStruct struct {
	Name     string
	IdNum    string
	Mobile   string
	PostCode string //邮编
	Province string //省份
	City     string //城市
	Region   string //区 县
	Street   string //街道
	Address  string //完整地址
}

func parsAddress(address string, blackAddress []string) bool {
	if len(blackAddress) == 0 {
		return true
	}
	for _, addr := range blackAddress {
		if strings.Contains(address, addr) {
			return false
		} else {
			if strings.Contains(addr, "-") {
				splitAddr := strings.Split(addr, "-")
				bnum := 0
				for _, v := range splitAddr {
					if strings.Contains(address, v) {
						bnum++
					}
				}
				if bnum == len(splitAddr) {
					return false
				}
			}
		}
	}
	return true
}

func resolutionAddress(address string) addrStruct {
	parse := Smart(address)
	addrStruct := addrStruct{}
	addrStruct.Name = parse.Name
	addrStruct.IdNum = parse.IdNumber
	addrStruct.Mobile = parse.Mobile
	addrStruct.PostCode = parse.PostCode
	addrStruct.Province = parse.Province
	addrStruct.City = parse.City
	addrStruct.Region = parse.Region
	addrStruct.Street = parse.Street
	addrStruct.Address = parse.Address
	return addrStruct
}

func TestSmart(t *testing.T) {
	userAddress := "上海市新疆南路秋屋小区"
	address := resolutionAddress(userAddress)
	fmt.Printf("%+v", address)
	//res := parsAddress(userAddress, blackAddress)
	//fmt.Println(res)
}
