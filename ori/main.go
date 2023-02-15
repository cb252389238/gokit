package main

import (
	"log"
	"ori/internal/core/ori"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	ori.Start() //启动项目
}
