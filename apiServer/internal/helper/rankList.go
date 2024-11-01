package helper

import (
	"time"
)

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
