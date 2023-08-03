package php

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"reflect"
	"time"
)

/*
多个数组切片合成一个数组切片
*/
func Array_merge[T comparable](a ...[]T) []T {
	if len(a) <= 0 {
		return []T{}
	}
	l := 0
	for _, arr := range a {
		l += len(arr)
	}
	res := make([]T, 0, l)
	for _, arr := range a {
		for _, v := range arr {
			res = append(res, v)
		}
	}
	return res
}

/*
截取切片数组中得一部分
*/
func Array_slice[T comparable](arr []T, start, length int) []T {
	res := []T{}
	if start > len(arr)-1 {
		return res
	}
	if start+length > len(arr) {
		return res
	}
	return arr[start : start+length]
}

/*
返回差集
*/

func Array_diff[T comparable](arr ...[]T) []T {
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

/*
返回交集
*/

func Array_intersect[T comparable](arr ...[]T) []T {
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

/*
检查键名是否在数组中
*/

func Array_key_exists[T comparable](key T, arr map[T]any) bool {
	if _, ok := arr[key]; ok {
		return true
	}
	return false
}

/*
检查val是否在数组中
*/
func In_array(needle any, haystack any) bool {
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

/*
获取数组元素数量
*/
func Count[T phpArray](arr T) int {
	return len(arr)
}

/*
*数组值去重
 */
func Array_unique[T comparable](arr []T) []T {
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

/*
从数组中随机抽出n个元素
*/
func Array_rand[T comparable](elements []T) []T {
	c, _ := crand.Int(crand.Reader, big.NewInt(99999))
	r := rand.New(rand.NewSource(time.Now().UnixNano() + c.Int64()))
	n := make([]T, len(elements))
	for i, v := range r.Perm(len(elements)) {
		n[i] = elements[v]
	}
	return n
}

/*
返回数组中所有键组成新得数组
*/
func Array_keys(elements map[any]any) []any {
	i, keys := 0, make([]any, len(elements))
	for key := range elements {
		keys[i] = key
		i++
	}
	return keys
}

/*
返回数组中所有得值
*/
func ArrayValues(elements map[any]any) []any {
	i, vals := 0, make([]any, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

/*
反转数组得键值对
*/
func Array_flip[T1 comparable, T2 comparable](m map[T1]T2) map[T2]T1 {
	n := make(map[T2]T1)
	for i, v := range m {
		n[v] = i
	}
	return n
}

/*
数组反转
*/
func Array_reverse[T comparable](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

/*
统计数组中元素出现得次数
*/
func Array_count_values[T comparable](s []T) map[T]int {
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

/*
打乱数组中得元素
*/
func Shuffle(array any) {
	valueOf := reflect.ValueOf(array)
	if valueOf.Type().Kind() != reflect.Slice {
		return
	}
	length := valueOf.Len()
	if length < 2 {
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	swapper := reflect.Swapper(array)
	for i := 0; i < length; i++ {
		j := r.Intn(length)
		swapper(i, j)
	}
}

/*
删除数组中得第一个元素
*/
func Array_shift[T comparable](array []T) []T {
	if len(array) == 0 {
		return array
	}
	return array[1:]
}

/*
删除数组中最后一个元素
*/
func Array_pop[T comparable](array []T) []T {
	if len(array) == 0 {
		return array
	}
	return array[0 : len(array)-1]
}

/*
在数组尾部插入元素
*/
func Array_push[T comparable](array []T, val ...T) []T {
	return append(array, val...)
}

/*
在数组开头插入元素
*/

func Array_unshift[T comparable](array *[]T, values ...T) int {
	*array = append(values, *array...)
	return len(*array)
}
