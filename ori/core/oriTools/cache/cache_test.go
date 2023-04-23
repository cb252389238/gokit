package cache

import (
	"fmt"
	"testing"
	"time"
)

type TestStruct struct {
	Num      int
	Children []*TestStruct
}

func TestSetGet(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 1, time.Second*5)
	fmt.Println(tc.Get("1"))
	time.Sleep(time.Second * 6)
	fmt.Println(tc.Get("1"))
}

func TestCache_Ttl(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 1, time.Second*5)
	fmt.Println(tc.Ttl("1"))
	time.Sleep(time.Second)
	fmt.Println(tc.Pttl("1"))
}

func TestCache_SetNx(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 1, time.Second*5)
	fmt.Println(tc.SetNx("1", 1, time.Second*5))
	time.Sleep(time.Second * 6)
	fmt.Println(tc.SetNx("1", 1, time.Second*5))
}

func TestCache_Increment(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 1.2, time.Second*5)
	err := tc.Increment("1", 1.1)
	fmt.Println(err)
	fmt.Println(tc.Get("1"))
}

func TestCache_Decrement(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 5, time.Second*5)
	err := tc.Decrement("1", 1)
	fmt.Println(err)
	fmt.Println(tc.Get("1"))
}

func TestCache_Delete(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 5, time.Second*5)
	tc.Delete("1")
	fmt.Println(tc.Get("1"))
}

func TestCache_SaveFile(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.Set("1", 5, time.Second*500)
	err := tc.SaveFile("data.txt")
	fmt.Println(err)
}

func TestCache_LoadFile(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.LoadFile("data.txt")
	fmt.Println(tc.get("1"))
}

func TestCache_OnEvicted(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.OnEvicted(func(s string, i any) {
		fmt.Printf("key:%s,v:%v\r\n", s, i)
	})
	tc.Set("1", 10, time.Second*5)
	time.Sleep(time.Second * 10)
}

func TestCache_Expire(t *testing.T) {
	tc := New(DefaultExpiration, time.Second*1)
	tc.OnEvicted(func(s string, i any) {
		fmt.Printf("key:%s,v:%v\r\n", s, i)
	})
	tc.Set("1", 1, time.Second*3)
	fmt.Println("设置值")
	time.Sleep(time.Second * 2)
	fmt.Println(tc.Expire("1", time.Second*3))
	fmt.Println("续命1")
	time.Sleep(time.Second * 2)
	fmt.Println(tc.Expire("1", time.Second*3))
	fmt.Println("续命2")
	time.Sleep(time.Second * 2)
	fmt.Println(tc.Expire("1", time.Second*3))
	fmt.Println("续命3")
	time.Sleep(time.Second * 2)
	fmt.Println(tc.Expire("1", time.Second*3))
	fmt.Println("续命4")
	time.Sleep(time.Second * 2)
	fmt.Println(tc.Expire("1", time.Second*3))
	fmt.Println("续命5")
	time.Sleep(time.Second * 10)
}
