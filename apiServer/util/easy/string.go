package easy

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

// JSONStringFormObject
//
//	@Description: 格式化为JSON字符串
//	@param object interface{} - 待格式化的值
//	@return string - JSON序列化值
func JSONStringFormObject(object interface{}) string {
	jsonObject := "{}"
	byteObject, err := json.Marshal(object)
	if err == nil {
		jsonObject = string(byteObject)
	}
	return jsonObject
}

// NumberToK
//
//	@Description: 任意数据类型数字格式化k单位字符串
//	@param src interface{} - 任意数据类型数字
//	@return string - 格式化k单位字符串 eg: 1234 - 1.23k
func NumberToK(src interface{}) string {
	num := cast.ToInt64(src)
	d := decimal.NewFromInt(num)
	if num > 999 {
		return d.Div(decimal.NewFromInt(1000)).StringFixed(2) + "k"
	}
	return cast.ToString(src)
}

// NumberToW
//
//	@Description: 任意数据类型数字格式化w单位字符串
//	@param src interface{} - 任意数据类型数字
//	@return string - 格式化w单位字符串 eg: 12345 - 1.23w
func NumberToW(src interface{}, fixed ...int32) string {
	num := cast.ToInt64(src)
	d := decimal.NewFromInt(num)
	def := int32(2)
	if len(fixed) > 0 {
		def = fixed[0]
	}
	if num > 9999999 {
		return d.Div(decimal.NewFromInt(10000000)).StringFixed(def) + "KW"
	}
	if num > 9999 {
		return d.Div(decimal.NewFromInt(10000)).StringFixed(def) + "W"
	}
	return cast.ToString(src)
}

// GetInSql
//
//	@Description: 格式化为in sql 语句
//	@param src []string - 目标数组
//	@return string - in sql 语句 eg: ("a","b","c")
func GetInSql(src []string) string {
	inSql := "("
	for index, item := range src {
		inSql += "'" + item + "'"
		if index != len(src)-1 {
			inSql += ","
		}
	}
	inSql += ")"
	return inSql
}

// GenLikeSql
//
//	@Description: 生成like模糊查询sql %src%
//	@param src string -
//	@return string -
func GenLikeSql(src string) string {
	return "%" + src + "%"
}

// versionCompare
//
//	@Description: 对比版本号大小
//	@param clientVersion 客户端版本号
//	@param serviceVersion 服务端版本号
//	@return bool true需要更新，false不需要更新
func VersionCompare(clientVersion, serviceVersion string) bool {
	v1 := strings.Split(clientVersion, ".")
	v2 := strings.Split(serviceVersion, ".")

	for i := 0; i < len(v1) || i < len(v2); i++ {
		num1 := 0
		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}

		num2 := 0
		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}

		if num1 == num2 {
			continue
		}
		if num1 < num2 {
			return true
		}
		if num1 > num2 {
			return false
		}
	}
	return false
}
