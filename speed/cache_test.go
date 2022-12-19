package speed

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSpeedKV(t *testing.T) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Set("key", 1, time.Second*5, false) //设置缓存 0为永不过期
	fmt.Println(c.SetNx("key", 1, 0, false))
	fmt.Println(c.Get("key"))   //获取结果 返回1 true
	fmt.Println(c.GetEx("key")) //获取结果以及过期时间 1 2022-07-30 14:35:19 +0800 CST true
	c.Del("key")                //删除缓存
	fmt.Println(c.Get("key"))   //<nil> false

	c.Set("key", 1, time.Second*5, false) //设置缓存 5秒后过期
	time.Sleep(time.Second * 3)           //生命周期还剩两秒
	c.Set("key", 1, time.Second*5, false) //再次设置相同的key，更新生命周期
	time.Sleep(time.Second * 3)           //生命周期还剩五秒
	fmt.Println(c.Get("key"))             //1 true

	//绑定回调函数，当主动删除缓存或者缓存过期触发  v就是设置的缓存值
	c.BindDeleteCallBackFunc(func(k string, v interface{}) {
		fmt.Println(v)
		fmt.Println("触发回调函数")
	})
	c.Set("test01", 100, time.Second*3, false)
	c.Del("test01")                           //callBack false 不触发回调函数
	c.Set("test02", 200, time.Second*2, true) //两秒后过期触发回调函数
	time.Sleep(time.Second * 10)              //阻塞
}

func TestSpeedHash(t *testing.T) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.BindDeleteCallBackFunc(func(k string, v interface{}) {
		fmt.Println("触发回调函数")
	})
	c.HSet("userinfo", "name", "城邦")
	c.HSet("userinfo", "age", 30)
	c.HSet("userinfo", "sex", "男")
	c.HMSet("userinfo", map[string]interface{}{
		"aaa": 111,
		"bbb": 222,
	})
	fmt.Println("HSetNX", c.HSetNx("userinfo", "ccc", 30))
	fmt.Println("HSetNX", c.HSetNx("userinfo", "age", 30))
	fmt.Println(c.HExists("userinfo", "age"))
	fmt.Println(c.HExists("userinfo1", "class"))
	fmt.Println(c.HGet("userinfo", "age", "name"))
	fmt.Println(c.HKeys("userinfo"))
	fmt.Println(c.HVAls("userinfo"))
	c.HDel("userinfo", "age", "sex")
	fmt.Println("过期前:", c.HGetAll("userinfo"))
	c.HSetEx("userinfo", time.Second*1, true)
	time.Sleep(time.Second * 2)
	fmt.Println("过期后:", c.HGetAll("userinfo"))
}

func TestSpeedSet(t *testing.T) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.BindDeleteCallBackFunc(func(k string, v interface{}) {
		fmt.Println(v)
		fmt.Println("触发回调函数")
	})
	c.SAdd("1070", time.Second*1, true, 1001, 1002, 1003)
	c.SAdd("1070", time.Second*10, true, 1001)
	time.Sleep(time.Second * 5)
	fmt.Println(c.SMembers("1070"))
	//fmt.Println(c.SISMembers("1070", 1001))
	//fmt.Println(c.SISMembers("1070", 1005))
	//fmt.Println(c.SCard("1070"))
	fmt.Println(c.SMembers("1070"))
	time.Sleep(time.Second * 4)
	//fmt.Println(c.SISMembers("1070", 1001))
	fmt.Println(c.SMembers("1070"))
	c.SAdd("1070", time.Second*3, true, 1008)
	fmt.Println(c.SMembers("1070"))
	time.Sleep(time.Second * 10)
	fmt.Println(c.SMembers("1070"))
}

func TestSpeedSetRem(t *testing.T) {
	cache, err := New()
	if err != nil {
		fmt.Println("实例化缓存出错", err)
		return
	}
	cache.BindDeleteCallBackFunc(func(key string, val interface{}) {
		fmt.Printf("回调函数 key:%s,val:%v\r\n", key, val)
	})
	cache.Set("name", "城邦", time.Second*5, true)
	if get, b := cache.Get("name"); b {
		fmt.Println("val:", get)
	}
	cache.SAdd("members", time.Second*10, false, 1001)
	cache.SAdd("members", time.Second*11, true, 1002)
	cache.SAdd("members", time.Second*12, true, 1003)
	cache.SAdd("members", time.Second*13, true, 1004)
	cache.SAdd("members", time.Second*14, true, 1005)
	fmt.Println(cache.SMembers("members"))
	fmt.Println(cache.SRem("members", 1001))
	fmt.Println(cache.SRem("members", 1002))
	fmt.Println(cache.SMembers("members"))
	select {}
}

func BenchmarkCache_SetEx(b *testing.B) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		c.Set("key-"+strconv.Itoa(i), i, time.Second*60, false)
	}
}

func BenchmarkCache_SetExAndDel(b *testing.B) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		c.Set("key-"+strconv.Itoa(i), i, time.Second*60, false)
	}
	for i := 0; i < b.N; i++ {
		c.Del("key-" + strconv.Itoa(i))
	}
}
func BenchmarkCache_Set(b *testing.B) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		c.Set("key-"+strconv.Itoa(i), i, 0, false)
	}
}

func BenchmarkCacheSetRepeat(b *testing.B) {
	c, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		c.Set("key", i, time.Second*60, false)
	}
}
