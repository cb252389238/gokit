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

func Strtotime(format, strtime string) (int64, error) {
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

func Checkdate(month, day, year int) bool {
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

func Sleep(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}

func Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
