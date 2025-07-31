package core_IpToRegionCode

import (
	"strings"

	"github.com/ip2location/ip2location-go"
)

var ipDb *ip2location.DB

func init() {
	ipDb, _ = ip2location.OpenDB("IP2LOCATION-LITE-DB1.BIN")
}

func IpToRegionCode(ip string) (regionCode string, err error) {
	countryCode, err := ipToCountryCode(ip)
	if err != nil {
		return
	}
	code := getCallingCode(countryCode)
	return code, nil
}

func ipToCountryCode(ip string) (code string, err error) {
	if ipDb == nil {
		ipDb, err = ip2location.OpenDB("IP2LOCATION-LITE-DB1.BIN")
		if err != nil {
			return "", err
		}
	}
	results, err := ipDb.Get_all(ip)
	if err != nil {
		return "", err
	}
	return results.Country_short, nil
}

func getCallingCode(countryCode string) string {
	// 转换为大写以统一处理
	countryCode = strings.ToUpper(countryCode)

	// 国家代码到国际电话区号的映射
	countryCallingCodes := map[string]string{
		// 亚洲
		"CN": "+86",  // 中国
		"JP": "+81",  // 日本
		"KR": "+82",  // 韩国
		"IN": "+91",  // 印度
		"SG": "+65",  // 新加坡
		"MY": "+60",  // 马来西亚
		"TH": "+66",  // 泰国
		"VN": "+84",  // 越南
		"ID": "+62",  // 印度尼西亚
		"PH": "+63",  // 菲律宾
		"HK": "+852", // 香港(中国)
		"MO": "+853", // 澳门(中国)
		"TW": "+886", // 台湾(中国)
		"IL": "+972", // 以色列
		"SA": "+966", // 沙特阿拉伯
		"AE": "+971", // 阿联酋
		"TR": "+90",  // 土耳其

		// 欧洲
		"RU": "+7",   // 俄罗斯
		"DE": "+49",  // 德国
		"FR": "+33",  // 法国
		"IT": "+39",  // 意大利
		"ES": "+34",  // 西班牙
		"GB": "+44",  // 英国
		"NL": "+31",  // 荷兰
		"BE": "+32",  // 比利时
		"CH": "+41",  // 瑞士
		"AT": "+43",  // 奥地利
		"SE": "+46",  // 瑞典
		"FI": "+358", // 芬兰
		"NO": "+47",  // 挪威
		"DK": "+45",  // 丹麦
		"PL": "+48",  // 波兰
		"UA": "+380", // 乌克兰

		// 北美洲
		"US": "+1",  // 美国
		"CA": "+1",  // 加拿大
		"MX": "+52", // 墨西哥

		// 南美洲
		"BR": "+55", // 巴西
		"AR": "+54", // 阿根廷
		"CO": "+57", // 哥伦比亚
		"PE": "+51", // 秘鲁
		"CL": "+56", // 智利

		// 大洋洲
		"AU": "+61", // 澳大利亚
		"NZ": "+64", // 新西兰

		// 非洲
		"ZA": "+27",  // 南非
		"EG": "+20",  // 埃及
		"NG": "+234", // 尼日利亚
		"KE": "+254", // 肯尼亚
	}

	if code, ok := countryCallingCodes[countryCode]; ok {
		return code
	}
	return "" // 未知国家代码返回空字符串
}
