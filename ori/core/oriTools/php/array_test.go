package php

import (
	"fmt"
	"testing"
)

func TestArray_merge(t *testing.T) {
	fmt.Println(Array_merge([]int{1, 2, 3}, []int{3, 4, 5}))
}

func TestArray_slice(t *testing.T) {
	fmt.Println(Array_slice([]int{0, 1, 2, 3, 4}, 0, 6))
}

func TestArray_diff(t *testing.T) {
	fmt.Println(Array_diff([]int{1, 2, 3}, []int{2, 3, 4}))
}

func TestArray_intersect(t *testing.T) {
	fmt.Println(Array_intersect([]int{1, 2, 3}, []int{2, 3, 4}))
}

func TestArray_key_exists(t *testing.T) {
	fmt.Println(Array_key_exists("name", map[string]any{"name": "22"}))
}

func TestIn_array(t *testing.T) {
	fmt.Println(In_array(123, map[string]int{"aa": 123}))
}

func TestCount(t *testing.T) {
	fmt.Println(Count(map[int]int{1: 1, 2: 2, 3: 3}))
}

func TestArray_unique(t *testing.T) {
	fmt.Println(Array_unique([]int{1, 1, 1, 1, 2, 3}))
}

func TestArray_rand(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(Array_rand([]int{1, 2, 3, 4, 5}))
	}
}

func TestArray_keys(t *testing.T) {
	fmt.Println(Array_keys(map[any]any{"1": 1, "2": 2}))
}

func TestArrayValues(t *testing.T) {
	fmt.Println(ArrayValues(map[any]any{1: 1, 2: 2}))
}

func TestArray_flip(t *testing.T) {
	fmt.Println(Array_flip(map[string]int{"1": 3, "2": 4}))
}

func TestArray_reverse(t *testing.T) {
	fmt.Println(Array_reverse([]int{1, 2, 3, 4}))
}

func TestArray_count_values(t *testing.T) {
	fmt.Println(Array_count_values([]int{1, 1, 2, 2, 3, 3, 3}))
}

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 3}
	Shuffle(a)
	fmt.Println(a)
}

func BenchmarkShuffle(b *testing.B) {
	a := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		a = append(a, i)
	}
	for i := 0; i < b.N; i++ {
		Shuffle(a)
	}
}

func TestArray_shift(t *testing.T) {
	fmt.Println(Array_shift([]int{1, 2, 3}))
}

func TestArray_pop(t *testing.T) {
	fmt.Println(Array_pop([]int{1, 2, 3}))
}

func TestArray_push(t *testing.T) {
	fmt.Println(Array_push([]int{1, 2, 3}, 4, 5))
}

func TestArray_unshift(t *testing.T) {
	a := []int{1, 2, 3}
	Array_unshift(&a, 4, 5)
	fmt.Println(a)
}
