package core

import (
	"github.com/spf13/cast"
	"testing"
)

func TestCheckConsecutive(t *testing.T) {
	ser := new(UserFancyNo)
	//测试连续4-8位重复数字
	ok, types := ser.checkConsecutive(cast.ToString(10012222))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(10022222))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(10222222))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(12222222))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(22222222))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(333300123))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(33303330))
	t.Log(ok, types, "false")
	ok, types = ser.checkConsecutive(cast.ToString(32555551))
	t.Log(ok, types, "true")
	ok, types = ser.checkConsecutive(cast.ToString(12312312))
	t.Log(ok, types, "false")
	ok, types = ser.checkConsecutive(cast.ToString(11122233))
	t.Log(ok, types, "false")
	ok, types = ser.checkConsecutive(cast.ToString(11112222))
	t.Log(ok, types, "true")
}

// ababab类型检测
func TestCheckABABAB(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkABABAB(cast.ToString(10121212))
	t.Log(ok, types, "true")
	ok, types = ser.checkABABAB(cast.ToString(12121234))
	t.Log(ok, types, "true")
	ok, types = ser.checkABABAB(cast.ToString(12341212))
	t.Log(ok, types, "false")
	ok, types = ser.checkABABAB(cast.ToString(90909090))
	t.Log(ok, types, "true")
	ok, types = ser.checkABABAB(cast.ToString(90909091))
	t.Log(ok, types, "true")
	ok, types = ser.checkABABAB(cast.ToString(53535335))
	t.Log(ok, types, "true")
	ok, types = ser.checkABABAB(cast.ToString(11121212))
	t.Log(ok, types, "true")
}

// 检测abbbabbb模式
func TestCheckABBBABBB(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkABBBABBB(cast.ToString(12221222))
	t.Log(ok, types, "true")
	ok, types = ser.checkABBBABBB(cast.ToString(31113111))
	t.Log(ok, types, "true")
	ok, types = ser.checkABBBABBB(cast.ToString(31123111))
	t.Log(ok, types, "false")
	ok, types = ser.checkABBBABBB(cast.ToString(33313331))
	t.Log(ok, types, "false")
	ok, types = ser.checkABBBABBB(cast.ToString(15555111))
	t.Log(ok, types, "false")
	ok, types = ser.checkABBBABBB(cast.ToString(12221444))
	t.Log(ok, types, "false")
}

// 检测bbbabbba模式
func TestCheckBBBABBBA(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkBBBABBBA(cast.ToString(22212221))
	t.Log(ok, types, "true")
	ok, types = ser.checkBBBABBBA(cast.ToString(11131113))
	t.Log(ok, types, "true")
	ok, types = ser.checkBBBABBBA(cast.ToString(11231113))
	t.Log(ok, types, "false")
	ok, types = ser.checkBBBABBBA(cast.ToString(13331113))
	t.Log(ok, types, "false")
	ok, types = ser.checkBBBABBBA(cast.ToString(55515552))
	t.Log(ok, types, "false")
	ok, types = ser.checkBBBABBBA(cast.ToString(66615552))
	t.Log(ok, types, "false")
}

// 连续递增或者递减
func TestCheckAscending(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkAscending(cast.ToString(12345670))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(76543212))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(23456781))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(98765432))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(13579246))
	t.Log(ok, types, "false")
	ok, types = ser.checkAscending(cast.ToString(11111111))
	t.Log(ok, types, "false")
	ok, types = ser.checkAscending(cast.ToString(91234567))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(76543210))
	t.Log(ok, types, "true")
	ok, types = ser.checkAscending(cast.ToString(96543210))
	t.Log(ok, types, "true")
}

// 规则5: 前5递增/递减 + 后3相同 (如12345555)
func TestCheckFront5IncDecTail3Same(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkFront5IncDecTail3Same(cast.ToString(12345555))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(12345556))
	t.Log(ok, types, "false")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(12345777))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(12345001))
	t.Log(ok, types, "false")

	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(54321111))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(54321001))
	t.Log(ok, types, "false")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(98765555))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3Same(cast.ToString(98765001))
	t.Log(ok, types, "false")
}

func TestCheckFront5IncDecTail3IncDec(t *testing.T) {
	ser := new(UserFancyNo)
	ok, types := ser.checkFront5IncDecTail3IncDec(cast.ToString(12345123))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(12345321))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(65432123))
	t.Log(ok, types, "true")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(98765321))
	t.Log(ok, types, "true")

	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(12345111))
	t.Log(ok, types, "false")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(12345541))
	t.Log(ok, types, "false")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(54321111))
	t.Log(ok, types, "false")
	ok, types = ser.checkFront5IncDecTail3IncDec(cast.ToString(97531135))
	t.Log(ok, types, "false")
}
