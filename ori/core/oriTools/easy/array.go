package easy

import (
	"math/rand"
	"reflect"
	"time"
)

// 合并数组切片
func ArrayMerge[T comparable](a ...[]T) []T {
	if len(a) <= 0 {
		return []T{}
	}
	l := 0
	for _, arr := range a {
		l += len(arr)
	}
	res := make([]T, 0, l)
	for _, arr := range a {
		res = append(res, arr...)
	}
	return res
}

// 截取切片数组中得一部分
func ArrayCut[T comparable](arr []T, start, length int) []T {
	res := []T{}
	if start > len(arr)-1 {
		return res
	}
	if start+length > len(arr) {
		return res
	}
	return arr[start : start+length]
}

// 返回切片差集
func ArrayDiff[T comparable](arr ...[]T) []T {
	res := []T{}
	if len(arr) <= 0 {
		return res
	}
	maps := map[T]int{}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			maps[v2] += 1
		}
	}
	for k, v := range maps {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// 获取切片交集
func ArrayIntersect[T comparable](arr ...[]T) []T {
	res := []T{}
	if len(arr) <= 0 {
		return res
	}
	maps := map[T]int{}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			maps[v2] += 1
		}
	}
	for k, v := range maps {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

// 检测元素是否在切片中
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
	default:
		return false
	}
	return false
}

// 去重
func ArrayUnique[T comparable](arr []T) []T {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	size := len(arr)
	result := make([]T, 0, size)
	temp := map[T]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; !ok {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// 从数组中随机抽出n个元素
func ArrayRand[T comparable](elements []T, n int) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999999)))
	res := make([]T, len(elements))
	for i, v := range r.Perm(len(elements)) {
		res[i] = elements[v]
	}
	return res[0:n]
}

// 返回数组中所有键组成新得数组
func ArrayKeys(elements map[any]any) []any {
	i, keys := 0, make([]any, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

// 返回数组中所有得值
func ArrayValues(elements map[any]any) []any {
	i, vals := 0, make([]any, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

// 反转数组得键值对
func ArrayFlip[T1 comparable, T2 comparable](m map[T1]T2) map[T2]T1 {
	n := make(map[T2]T1)
	for i, v := range m {
		n[v] = i
	}
	return n
}

// 数组反转
func ArrayReverse[T comparable](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 统计数组中元素出现得次数
func ArrayCountValues[T comparable](s []T) map[T]int {
	r := make(map[T]int)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
		} else {
			r[v] = 1
		}
	}
	return r
}

// 打乱数组中得元素
func Shuffle(array any) {
	valueOf := reflect.ValueOf(array)
	if valueOf.Type().Kind() != reflect.Slice {
		return
	}
	length := valueOf.Len()
	if length < 2 {
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63n(9999999)))
	swapper := reflect.Swapper(array)
	for i := 0; i < length; i++ {
		j := r.Intn(length)
		swapper(i, j)
	}
}
