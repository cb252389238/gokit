package php

import (
	"strings"
	"time"
)

func Time() int64 {
	return time.Now().Unix()
}

var formatMap = map[string]string{
	"Y": "2006",
	"m": "01",
	"d": "02",
	"H": "15",
	"i": "04",
	"s": "05",
}

func StrToTime(format, strtime string) (int64, error) {
	for k, v := range formatMap {
		format = strings.Replace(format, k, v, -1)
	}
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func Date(format string, timestamp ...int64) string {
	var t int64
	if len(timestamp) == 0 {
		t = time.Now().Unix()
	} else {
		t = timestamp[0]
	}
	for k, v := range formatMap {
		format = strings.Replace(format, k, v, -1)
	}
	return time.Unix(t, 0).Format(format)
}

func CheckDate(month, day, year int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}

	return true
}

// 睡眠指定秒数
func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

// 睡眠微秒
func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

// 获取到本周周日时间戳
func GetUntilSundayTime() int64 {
	now := time.Now()
	weekday := now.Weekday()
	intWeekday := int(weekday)
	if intWeekday == 0 {
		intWeekday = 7
	}
	sunday := now.AddDate(0, 0, 7-intWeekday)
	endOfDay := time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, 999999999, sunday.Location())
	duration := endOfDay.Sub(now)
	return int64(duration.Seconds())
}

// 获取到月底时间戳
func GetUntilMonthTime() int64 {
	now := time.Now()
	nextMonth := now.AddDate(0, 1, 0)
	endOfMonth := time.Date(nextMonth.Year(), nextMonth.Month(), 0, 23, 59, 59, 999999999, now.Location())
	duration := endOfMonth.Sub(now)
	seconds := int64(duration.Seconds())
	return seconds
}

// 获取距离今天结束时间戳
func GetUntilDayTime() int64 {
	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())
	duration := endOfDay.Sub(now)
	return int64(duration.Seconds())
}

// 获取本周周一日期
func GetMondayDateTime() string {
	now := time.Now()
	weekday := now.Weekday()
	intWeekday := int(weekday)
	if intWeekday == 0 {
		intWeekday = 7
	}
	monday := now.AddDate(0, 0, -intWeekday+1)
	return monday.Format("20060102")
}

// 获取本周周日日期
func GetSundayDateTime() string {
	now := time.Now()
	weekday := now.Weekday()
	intWeekday := int(weekday)
	if intWeekday == 0 {
		intWeekday = 7
	}
	sunday := now.AddDate(0, 0, 7-intWeekday)
	return sunday.Format("20060102")
}

// 获取上周周一日期
func GetLastMondayDateTime() string {
	now := time.Now()
	weekday := now.Weekday()
	intWeekday := int(weekday)
	if intWeekday == 0 {
		intWeekday = 7
	}
	lastMonday := now.AddDate(0, 0, -intWeekday-6)
	return lastMonday.Format("20060102")
}

// 获取上周周日日期
func GetLastSundayDateTime() string {
	now := time.Now()
	weekday := now.Weekday()
	intWeekday := int(weekday)
	if intWeekday == 0 {
		intWeekday = 7
	}
	lastSunday := now.AddDate(0, 0, -intWeekday)
	return lastSunday.Format("20060102")
}

// 获取本月月份
func GetMonthDateTime() string {
	now := time.Now()
	return now.Format("01")
}

// 获取上个月月份
func GetLastMonthDateTime() string {
	now := time.Now()
	return now.AddDate(0, -1, 0).Format("01")
}
