package helper

import (
	"fmt"
	"testing"
)

func TestGetMondayDateTime(t *testing.T) {
	t.Log(GetMondayDateTime())
}

func TestGetLastMondayDateTime(t *testing.T) {
	t.Log(GetLastMondayDateTime())
}

func TestGetSundayDateTime(t *testing.T) {
	t.Log(GetSundayDateTime())
}

func TestGetLastSundayDateTime(t *testing.T) {
	t.Log(GetLastSundayDateTime())
}

func TestGetMonthDateTime(t *testing.T) {
	t.Log(GetMonthDateTime())
}

func TestGetLastMonthDateTime(t *testing.T) {
	t.Log(GetLastMonthDateTime())
}

func TestGetUntilDayTime(t *testing.T) {
	fmt.Println(GetUntilDayTime())
}

func TestGetUntilSundayTime(t *testing.T) {
	fmt.Println(GetUntilSundayTime())
}

func TestGetUntilMonthTime(t *testing.T) {
	fmt.Println(GetUntilMonthTime())
}
