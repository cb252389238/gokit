package easy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// 获取本地ip
func GetLocalIp() string {
	var err error
	var ip = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	return ip
}

// ip转int类型
func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// int转string
func IntIpToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

func Ipv4IntToString(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func Ipv4StringToInt(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// 判断 IP 是否为内网 IP
func IsPrivateIP(ip string) bool {
	privateIPBlocks := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, block := range privateIPBlocks {
		_, ipNet, _ := net.ParseCIDR(block)
		if ipNet.Contains(net.ParseIP(ip)) {
			return true
		}
	}
	return false
}

// 通过第三方 API 获取公网 IP 的具体地址
func GetIPLocation(ip string) (address string, err error) {

	url := fmt.Sprintf("https://api.ipbase.com/v1/json/%s", ip) // 使用 免费服务
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	if _, ok := result["region_name"]; ok {
		address = result["region_name"].(string)
	}
	if _, ok := result["city"]; ok {
		address += "-" + result["city"].(string)
	}
	return
}
