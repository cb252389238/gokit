package main

import (
	"experiment/redis"
	"fmt"
)

func main() {
	r := redis.New(redis.WithHost("127.0.0.1"), redis.WithPort(6379))
	fmt.Println(r)
}
