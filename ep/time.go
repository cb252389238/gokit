package ep

import (
	"fmt"
	"strings"
	"time"
)

const (
	LAYOUT_YMD = "2006-01-02"
	LAYOUT_HMS = "15:04:05"
)

/**
字符串转时间戳
字符串格式：2022-03-25 14:42:30，2022-03-25 14-42-30
20220325 144230 中间没有间隔符也可以
2022-03-25，20220325 只有日期

*/
func StrToTime(timeStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	var time_layout string
	switch len(timeStr) {
	case 19:
		time_layout = LAYOUT_YMD + " " + LAYOUT_HMS
	case 15:
		time_layout = strings.Replace(LAYOUT_YMD, "-", "", -1) + strings.Replace(LAYOUT_HMS, ":", "", -1)
	case 10:
		time_layout = LAYOUT_YMD
	case 8:
		time_layout = strings.Replace(LAYOUT_YMD, "-", "", -1)
	}
	theTime, _ := time.ParseInLocation(time_layout, timeStr, loc)
	return theTime.Unix()
}

/**
时间戳转字符串
timeStamp 时间戳
layout Y-m-d H:i:s,Y-m-d,Ymd 等等
*/
func TimeToStr(timeStamp int64, layout string) string {
	if strings.Contains(layout, "Y") {
		layout = strings.Replace(layout, "Y", "2006", -1)
	}
	if strings.Contains(layout, "y") {
		layout = strings.Replace(layout, "y", "06", -1)
	}
	if strings.Contains(layout, "m") {
		layout = strings.Replace(layout, "m", "01", -1)
	}
	if strings.Contains(layout, "d") {
		layout = strings.Replace(layout, "d", "02", -1)
	}
	if strings.Contains(layout, "H") {
		layout = strings.Replace(layout, "H", "15", -1)
	}
	if strings.Contains(layout, "i") {
		layout = strings.Replace(layout, "i", "04", -1)
	}
	if strings.Contains(layout, "s") {
		layout = strings.Replace(layout, "s", "05", -1)
	}
	if timeStamp == 0 {
		return time.Now().Format(layout)
	} else {
		return time.Unix(timeStamp, 0).Format(layout)
	}
}

/**
获取时间戳秒
*/
func Time() int64 {
	return time.Now().Unix()
}

/**
获取时间戳纳秒
*/
func TimeNano() int64 {
	return time.Now().UnixNano()
}

/**
获取n天前后时间戳 秒
-1一天前
1 一天后
*/
func TimeDay(day int) int64 {
	return time.Now().AddDate(0, 0, day).Unix()
}

/**
获取n天前后0点时间戳
-1一天前
1 一天后
*/
func TimeDayStart(day int) int64 {
	timeStr := time.Now().Format(LAYOUT_YMD)
	t, _ := time.ParseInLocation(LAYOUT_YMD, timeStr, time.Local)
	return t.AddDate(0, 0, day).Unix()
}

func getLastDay() int {
	d := time.Now()
	d = d.AddDate(0, 0, -d.Day()+1)                                       //获取当月第一天
	d = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()) //时分秒重置为0
	fmt.Println(d.Year(), int(d.Month()), d.Day(), d.Hour())
	d = d.AddDate(0, 1, -1) //月份加1到下个月，然后日期减一到最后一天
	return d.Day()
}
