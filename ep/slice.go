package ep

import "reflect"

func IsExistItem(slice interface{}, value interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func InArray[T comparable](val T, arr []T) bool {
	ok := false
	for _, v := range arr {
		if v == val {
			ok = true
			break
		}
	}
	return ok
}
