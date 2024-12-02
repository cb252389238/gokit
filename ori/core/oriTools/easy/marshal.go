package easy

import (
	"reflect"
	"strconv"
	"strings"
)

type marshal struct {
}

// json反序列化
func Marshal(obj interface{}) string {
	v := reflect.ValueOf(obj)
	return new(marshal).parse(v)
}

// 处理开始
func (this *marshal) parse(v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.String:
		return `"` + v.String() + `"`
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Struct:
		return this.parseStruct(v, v.Type())
	case reflect.Map:
		return this.parseMap(v, v.Type())
	case reflect.Slice:
		return this.parseSlice(v, v.Type())
	case reflect.Interface:
		return this.parse(reflect.ValueOf(v.Interface()))
	default:

	}
	return "null"
}

// 处理结构体
func (this *marshal) parseStruct(v reflect.Value, t reflect.Type) string {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var str = make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fieldName := field.Name
		tag := field.Tag.Get("json")
		if tag != "" {
			f := strings.Split(tag, ",")
			fieldName = f[0]
		}
		if fieldName == "-" {
			continue
		}
		str[i] = `"` + fieldName + `":` + this.parse(value)
	}
	return "{" + strings.Join(str, ",") + "}"
}

// 处理切片
func (this *marshal) parseSlice(v reflect.Value, t reflect.Type) string {
	var str = make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		str[i] = this.parse(v.Index(i))
	}
	return `[` + strings.Join(str, ",") + "]"
}

// 处理map
func (this *marshal) parseMap(v reflect.Value, t reflect.Type) string {
	var str = make([]string, v.Len())
	m := v.MapRange()
	var i int
	for m.Next() {
		str[i] = `"` + m.Key().String() + `":` + this.parse(m.Value())
		i++
	}
	return "{" + strings.Join(str, ",") + "}"
}
