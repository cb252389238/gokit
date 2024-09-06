gokit
======

![](https://img.shields.io/badge/golang-1.20+-blue.svg?style=flat)
[![unit test](https://github.com/awa/go-iap/actions/workflows/unit_test.yml/badge.svg)](https://github.com/awa/go-iap/actions/workflows/unit_test.yml)

>Gokit is a scaffolding toolkit that you can install via "go get", or download and modify it locally to suit your specific needs.

# Package introduction:

* addrParse:  A tool to parse address.
* chscht: Convert between simplified and traditional Chinese characters.
* cronParse:  A tool to parse crontab.
* goAuth: A universal role-based access control system implemented in Go.
* idcard: Chinese ID card recognition tool, which identifiesChinese ID card recognition tool, which identifies date of birth and other information which identifies the region, gender, date of birth and other information through the ID number.
* imMerge: Automatically merge multiple messages for use in websocket push scenarios, combining messages to reduce the frequency of pushes.
* msgStore: Message caching can be applied to message retrieval services.
* ori: This is a web project scaffolding, which includes a variety of commonly used tools.
* projectRows: A tool for counting the number of lines of code in a project.
* xiezhi: Document similarity comparison, including algorithmsDocument similarity comparison, including algorithms distance, Simhash, including algorithms such as Hamming distance, Simhash, Minhash, and Cosine similarity.


# Installation
```
go get https://github.com/cb252389238/gokit
or
Clone to local
```


# Quick Start

### addrParse

```go

package addr

import (
  "fmt"
  "strings"
  "testing"
)

var blackAddress = []string{"福建-泉州", "福建-漳州", "厦门-思明", "新疆", "西藏", "吉林", "辽宁"}

type addrStruct struct {
  Name     string
  IdNum    string
  Mobile   string
  PostCode string //邮编
  Province string //省份
  City     string //城市
  Region   string //区 县
  Street   string //街道
  Address  string //完整地址
}

func parsAddress(address string, blackAddress []string) bool {
  if len(blackAddress) == 0 {
    return true
  }
  for _, addr := range blackAddress {
    if strings.Contains(address, addr) {
      return false
    } else {
      if strings.Contains(addr, "-") {
        splitAddr := strings.Split(addr, "-")
        bnum := 0
        for _, v := range splitAddr {
          if strings.Contains(address, v) {
            bnum++
          }
        }
        if bnum == len(splitAddr) {
          return false
        }
      }
    }
  }
  return true
}

func resolutionAddress(address string) addrStruct {
  parse := Smart(address)
  addrStruct := addrStruct{}
  addrStruct.Name = parse.Name
  addrStruct.IdNum = parse.IdNumber
  addrStruct.Mobile = parse.Mobile
  addrStruct.PostCode = parse.PostCode
  addrStruct.Province = parse.Province
  addrStruct.City = parse.City
  addrStruct.Region = parse.Region
  addrStruct.Street = parse.Street
  addrStruct.Address = parse.Address
  return addrStruct
}

func TestSmart(t *testing.T) {
  userAddress := "上海市新疆南路秋屋小区"
  address := resolutionAddress(userAddress)
  fmt.Printf("%+v", address)
  //res := parsAddress(userAddress, blackAddress)
  //fmt.Println(res)
}

```


### chscht

```go
func main() {
  dicter, err := New()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(dicter.Traditional("中华人民共和国"))
}
```

### cronParse

```go
func TestParse(t *testing.T) {
  expr, err := Parse("*/10 * * * * * * ")
  if err != nil {
    fmt.Println(err)
  }
  now := time.Now()
  next := expr.Next(now)
  fmt.Println(now.Unix())
  fmt.Println(next.Unix())
}
```

### goAuth

```go
package coreAuth

import (
  "fmt"
  "testing"
)

func TestGoAuth_AddRule(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  if err := goAuth.AddRule("enter_room_hidden", "隐身进厅", 1, "开启后隐身进入房间大厅"); err != nil {
    fmt.Println(err)
  }
  if err := goAuth.AddRule("up_maiwei", "上麦位", 1, "主动上麦位"); err != nil {
    fmt.Println(err)
  }
  if err := goAuth.AddRule("withdraw", "提现", 2, "提现权限"); err != nil {
    fmt.Println(err)
  }
}

func TestGoAuth_EditRule(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  if err := goAuth.EditRule(1, "enter_room_hidden", "隐身进入房间", 1, ""); err != nil {
    fmt.Println(err)
  }
}

func TestCoreAuth_DeleteRule(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  goAuth.DeleteRule(2)
}

func TestGoAuth_AddRole(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  if err := goAuth.AddRole("主持人", "1,3"); err != nil {
    t.Error(err)
  }
  if err := goAuth.AddRole("麦未嘉宾", "1"); err != nil {
    t.Error(err)
  }
  if err := goAuth.AddRole("歌手", "1"); err != nil {
    t.Error(err)
  }
}

func TestGoAuth_GiveUserRole(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  if err := goAuth.GiveUserRole(1, 1); err != nil {
    t.Error(err)
  }
  if err := goAuth.GiveUserRole(1, 2); err != nil {
    t.Error(err)
  }
}

func TestGoAuth_ShowRoleList(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  list, err := goAuth.ShowRoleList()
  if err != nil {
    t.Error(err)
  }
  fmt.Println(list)
}

func TestGoAuth_GetRoleRules(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  list, err := goAuth.GetRoleRules(5)
  if err != nil {
    t.Error(err)
  }
  fmt.Println(list)
}

func TestCoreAuth_VerifyAuth(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  auth, err := goAuth.VerifyAuth(2, "withdraw")
  fmt.Println(auth, err)
}

func TestCoreAuth_GetUserRules(t *testing.T) {
  goAuth, err := New(CoreAuthConfig{
    UserName: "root",
    PassWord: "root",
    Host:     "127.0.0.1",
    Port:     3306,
    Database: "test",
  })
  if err != nil {
    panic(err)
  }
  rules, err := goAuth.GetUserRules(2, 2)
  fmt.Println(rules, err)
}

```

### idcard

```go

func GetCitizenNoInfo(citizenNo []byte) (birthday int, err error) {
err = nil
birthday = 0
//isMale = false
//addrMask = 0
if !IsValidCitizenNo(&citizenNo) {
err = errors.New("Invalid citizen number.")
return
}

// Birthday information.
birthday, _ = strconv.Atoi(string(citizenNo[6:10]))
nMonth, _ := strconv.Atoi(string(citizenNo[10:12]))
nDay, _ := strconv.Atoi(string(citizenNo[12:14]))
birthday = time.Date(nYear, time.Month(nMonth), nDay, 0, 0, 0, 0, time.Local).Unix()

// Gender information.
genderMask, _ := strconv.Atoi(string(citizenNo[16]))
if genderMask%2 == 0 {
	isMale = false
} else {
	isMale = true
}

// Address code mask.
addrMask, _ = strconv.Atoi(string(citizenNo[:2]))

    return
}

```

### xiezhi

```go
package xiezhi

import (
  "strings"
  "sync"
  "xiezhi/cosinesim"
  "xiezhi/fenci"
  "xiezhi/jaccard"
  "xiezhi/minhash"
  "xiezhi/simhash"
  "xiezhi/util/charchar"
)

var (
  once sync.Once
  gose *fenci.FenCi
)

func NewFenCi() *fenci.FenCi {
  once.Do(func() {
    x := fenci.NewFenCi()
    gose = x
  })
  return gose
}

// 分词
func cut(text string) []string {
  text = charchar.RemovePunct(text)        //去除标点符号
  text = charchar.RemoveNonsenseWord(text) //去除无意义词语
  text = strings.Replace(text, " ", "", -1)
  text = strings.Replace(text, "\t", "", -1)
  text = strings.Replace(text, "\n", "", -1)
  text = strings.Replace(text, "\r", "", -1)
  return NewFenCi().Cut(text)
}

// 获取文档min hash签名
func SimHash(text string) uint64 {
  hash := simhash.SimHash(cut(text))
  return hash
}

// 对比两个文档hash相似性
// 返回海明距离和相似性
func SimHashSimilarity(hash1, hash2 uint64) (int, float64) {
  return simhash.Similarity(hash1, hash2)
}

// 获取hash签名
func MinHash(text string) []uint32 {
  return minhash.ComputeMinHashSignature(cut(text))
}

func MinHashSimilarity(hash1, hash2 []uint32) float64 {
  return minhash.ComputeSimilarity(hash1, hash2)
}

// 杰卡德系数计算
func Jaccard(text1, text2 string) float64 {
  coefficient := jaccard.ComputeJaccardCoefficient(cut(text1), cut(text2))
  return coefficient
}

// 余弦相似度计算
func CosineSim(text1, text2 string) float64 {
  frequency1 := cosinesim.CalculateTermFrequency(cut(text1))
  frequency2 := cosinesim.CalculateTermFrequency(cut(text2))
  similarity := cosinesim.ComputeCosineSimilarity(frequency1, frequency2)
  return similarity
}

```