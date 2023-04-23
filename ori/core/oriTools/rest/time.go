package rest

import "time"

const (
	LAYOUT_YMD = "2006-01-02"
	LAYOUT_HMS = "15:04:05"
)

func TimeDayStart(day int) int64 {
	timeStr := time.Now().Format(LAYOUT_YMD)
	t, _ := time.ParseInLocation(LAYOUT_YMD, timeStr, time.Local)
	return t.AddDate(0, 0, day).Unix()
}
