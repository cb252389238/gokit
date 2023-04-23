package php

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	ip := "127.0.0.1"
	long := Ip2long(ip)
	fmt.Println(long)
	long2ip := Long2ip(long)
	fmt.Println(long2ip)
}
