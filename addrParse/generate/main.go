package main

import "addrParse/generate/autoCode"

//go:generate go run main.go
func main() {
	autoCode.AutoAreaMap()
}
