package easy

import (
	"reflect"
)

// InSlice
//
//	@Description: 判断元素是否在切片中
//	@param needle
//	@param haystack
//	@return bool
func InArray(needle any, haystack any) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	}
	return false
}
