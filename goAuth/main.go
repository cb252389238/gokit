package main

import (
	"fmt"
	"goAuth/auth"
)

func main() {
	goAuth, err := auth.New("root:root@tcp(127.0.0.1:3306)/中文情感分析?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = goAuth.GiveUserRole(3, 2)
	fmt.Println(err)
}
