package easy

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// GetCurrMonthStartTime
//
//	@Description: 获取当前时间的月份的开始时间 00:00:00
//	@param now time.Time - 当前时间
//	@return time.Time - 当前时间月初的零点时间 00:00:00
func GetCurrMonthStartTime(now time.Time) time.Time {
	// 获取当前时间的年、月和日
	year, month, _ := now.Date()
	// 获取本月的第一天
	return time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
}

// GetCurrMonthEndTime
//
//	@Description: 获取当前时间的月份的开始时间 23:59:59
//	@param now time.Time - 当前时间
//	@return time.Time -  当前时间月末的时间 23:59:59
func GetCurrMonthEndTime(now time.Time) time.Time {
	// 提取当前时间的年份和月份
	year, month, _ := now.Date()
	// 获取下一个月的第一天
	nextMonthFirstDay := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())
	// 计算当前月份的最后一天的23:59:59
	return nextMonthFirstDay.Add(-1 * time.Second)
}

// GetCurrWeekStartTime
//
//	@Description: 获取当前时间的周一零点 00:00:00
//	@param now time.Time - 当前时间
//	@return time.Time - 当前时间周一的零点 00:00:00
func GetCurrWeekStartTime(now time.Time) time.Time {
	// 计算本周一的时间
	// 获取当前时间是周几
	weekday := now.Weekday()
	// 根据今天是周几计算偏移量
	var offset int
	if weekday == time.Sunday {
		// 如果今天是周日，则偏移量为-6，因为周日是一周的最后一天
		offset = -6
	} else {
		// 其他情况，偏移量为当前周几减去周一的差值
		offset = -int(weekday - time.Monday)
	}
	// 获取本周一的时间
	monday := now.AddDate(0, 0, offset)
	// 将本周一的时间重置为当天的零点
	return time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, monday.Location())
}

// GetCurrWeekEndTime
//
//	@Description: 获取当前时间的周末结束时间 23:59:59
//	@param now time.Time - 当前时间
//	@return time.Time - 当前时间周末的结束时间 23:59:59
func GetCurrWeekEndTime(now time.Time) time.Time {
	// 计算本周末的时间
	// 获取当前时间是周几
	weekday := now.Weekday()
	// 根据今天是周几计算偏移量
	var offset = 7
	if weekday == time.Sunday {
		// 如果今天是周日，则偏移量为0
		offset = 0
	} else {
		// 其他情况，偏移量为当前周几减去周一的差值
		offset = offset - int(weekday)
	}
	// 获取本周末的时间
	sunday := now.AddDate(0, 0, offset)
	// 将本周末的时间重置为当天的零点
	return time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, 0, sunday.Location())
}

// GetCurrDayStartTime
//
//	@Description: 获取当前时间零点 00:00:00
//	@param now time.Time - 当前时间
//	@return time.Time -  当前时间的零点 00:00:00
func GetCurrDayStartTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// GetCurrDayEndTime
//
//	@Description: 获取当前时间一天的结束时间 23:59:59
//	@param now time.Time - 当前时间
//	@return time.Time -  当前时间一天的结束时间 23:59:59
func GetCurrDayEndTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}

// GetCurrentYearWeek
//
//	@Description: 获取当前年周
//	@return string - eg: 202425 - 2024年第25周
func GetCurrentYearWeek() string {
	year, week := time.Now().ISOWeek()
	return fmt.Sprintf("%v%v", year, week)
}

// GetCurrentYMD
//
//	@Description: 获取当前时间年月日(默认2006-01-02)
//	@param now time.Time - 当前时间
//	@param format ...string - 格式化字符串
//	@return string - 当前时间年月日(默认2006-01-02)
func GetCurrentYMD(now time.Time, format ...string) string {
	defaultFormat := time.DateOnly
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return now.Format(defaultFormat)
}

// GetCurrentYMDHMS
//
//	@Description: 获取当前时间年月日时分秒(默认2006-01-02 15:04:05)
//	@param now time.Time - 当前时间
//	@param format ...string - 格式化字符串
//	@return string - 当前时间年月日时分秒(默认2006-01-02 15:04:05)
func GetCurrentYMDHMS(now time.Time, format ...string) string {
	defaultFormat := time.DateTime
	if len(format) > 0 {
		defaultFormat = format[0]
	}
	return now.Format(defaultFormat)
}

// SecondFormatString
//
//	@Description: 秒数格式化为时间格式 hh:mm:ss
//	@param seconds int64 - 秒数
//	@return string - 格式化的时间 hh:mm:ss
func SecondFormatString(seconds int64) string {
	defaultFormat := "%02d:%02d:%02d"
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf(defaultFormat, hours, minutes, remainingSeconds)
}

// 1. 解决gorm查询出的时间带时区问题
type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	if tTime.IsZero() {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.IsZero() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 判断是否是今天
func NowTimeIsToday(t time.Time) bool {
	now := time.Now()
	return t.Format(time.DateOnly) == now.Format(time.DateOnly)
}

// 判断是否是本周
func NowTimeIsThisWeek(t time.Time) bool {
	now := time.Now()
	_, thisWeek := now.ISOWeek()
	_, tWeek := t.ISOWeek()
	return now.Year() == t.Year() && thisWeek == tWeek
}

// 判断是否是本月
func NowTimeIsThisMonth(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() && now.Month() == t.Month()
}

// 获取本周的所有日期
func GetWeekDaysByTime(t time.Time) []string {
	// 获取本周第一天（周一）
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7 // 周日为0，需特殊处理
	}
	startOfWeek := t.AddDate(0, 0, -weekday+1)

	// 返回一周的七天
	var weekDays []string
	for i := 0; i < 7; i++ {
		weekDays = append(weekDays, startOfWeek.AddDate(0, 0, i).Format(time.DateOnly))
	}
	return weekDays
}

// 获取本月的所有日期
func GetMonthDaysByTime(t time.Time) []string {
	year, month, _ := t.Date()
	loc := t.Location()

	// 获取本月的第一天和下个月的第一天
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	nextMonth := firstOfMonth.AddDate(0, 1, 0)

	// 遍历本月的每一天
	var monthDays []string
	for day := firstOfMonth; day.Before(nextMonth); day = day.AddDate(0, 0, 1) {
		monthDays = append(monthDays, day.Format(time.DateOnly))
	}
	return monthDays
}
