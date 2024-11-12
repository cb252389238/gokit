package easy

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(Time())
}

func TestStrtotime(t *testing.T) {
	fmt.Println(StrToTime("Ymd H:i:s", "20230404 14:30:30"))
}

func TestDate(t *testing.T) {
	fmt.Println(Date("Y-m-d H:i:s", 1680618630))
}

func TestSleep(t *testing.T) {
	fmt.Println("start")
	Sleep(1)
	fmt.Println("end")
}

func TestUsleep(t *testing.T) {
	fmt.Println("start")
	Usleep(1000000)
	fmt.Println("end")
}

func TestCheckdate(t *testing.T) {
	fmt.Println(CheckDate(12, 02, 2023))
}
