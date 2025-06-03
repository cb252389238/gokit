package tools

import "time"

// 获取两个日期之间的天数差
func GetDaysBetweenDates(startDate, endDate time.Time) int {
	// 计算两个日期之间的天数差
	days := int(endDate.Sub(startDate).Hours() / 24)
	return days
}

// 获取本周开始date和结束date
func GetWeekStartAndEndTime(t time.Time) (string, string) {
	now := t
	weekStart := now.AddDate(0, 0, -int(now.Weekday())+1)
	if now.Weekday() == time.Sunday {
		weekStart = now.AddDate(0, 0, -6)
	}
	weekStart = time.Date(weekStart.Year(), weekStart.Month(), weekStart.Day(), 0, 0, 0, 0, weekStart.Location())
	weekEnd := weekStart.AddDate(0, 0, 7).Add(-time.Nanosecond)
	return weekStart.Format("2006-01-02"), weekEnd.Format("2006-01-02")
}
