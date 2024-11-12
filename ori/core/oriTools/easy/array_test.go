package easy

import (
	"fmt"
	"reflect"
	"testing"
)

type ComparableSlice[T comparable] struct {
	name     string
	arrays   [][]T
	expected []T
}

func TestArrayMerge(t *testing.T) {
	test1 := ComparableSlice[int]{
		name:     "merge multiple arrays of integers",
		arrays:   [][]int{{1, 2}, {3, 4}, {5}},
		expected: []int{1, 2, 3, 4, 5},
	}
	t.Run(test1.name, func(t *testing.T) {
		got := ArrayMerge(test1.arrays...)
		if !reflect.DeepEqual(got, test1.expected) {
			t.Errorf("ArrayMerge() = %v, want %v", got, test1.expected)
		}
	})
	test2 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"3", "4"}},
		expected: []string{"1", "2", "3", "4"},
	}
	t.Run(test2.name, func(t *testing.T) {
		got := ArrayMerge(test2.arrays...)
		if !reflect.DeepEqual(got, test2.expected) {
			t.Errorf("ArrayMerge() = %v, want %v", got, test2.expected)
		}
	})

	test3 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"", "4"}, {}},
		expected: []string{"1", "2", "", "4"},
	}
	t.Run(test3.name, func(t *testing.T) {
		got := ArrayMerge(test3.arrays...)
		if !reflect.DeepEqual(got, test3.expected) {
			t.Errorf("ArrayMerge() = %v, want %v", got, test3.expected)
		}
	})
}

func TestArrayCut(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		start    int
		length   int
		expected []int
	}{
		{"cut from middle", []int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4}},
		{"cut from start", []int{1, 2, 3, 4, 5}, 0, 2, []int{1, 2}},
		{"cut from end", []int{1, 2, 3, 4, 5}, 3, 2, []int{4, 5}},
		{"start out of range", []int{1, 2, 3, 4, 5}, 5, 1, []int{}},
		{"length out of range", []int{1, 2, 3, 4, 5}, 2, 3, []int{3, 4, 5}},
		{"empty array", []int{}, 0, 1, []int{}},
		{"complete array", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ArrayCut[int](tt.arr, tt.start, tt.length)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("ArrayCut(%v, %d, %d) = %v; expected %v", tt.arr, tt.start, tt.length, actual, tt.expected)
			}
		})
	}
}

func TestArrayDiff(t *testing.T) {
	test1 := ComparableSlice[int]{
		name:     "merge multiple arrays of integers",
		arrays:   [][]int{{1, 2}, {3, 4}, {5}},
		expected: []int{1, 2, 3, 4, 5},
	}
	t.Run(test1.name, func(t *testing.T) {
		got := ArrayDiff(test1.arrays...)
		if !reflect.DeepEqual(got, test1.expected) {
			t.Errorf("ArrayDiff() = %v, want %v", got, test1.expected)
		}
	})
	test2 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"1", "2", "3"}},
		expected: []string{"3"},
	}
	t.Run(test2.name, func(t *testing.T) {
		got := ArrayDiff(test2.arrays...)
		if !reflect.DeepEqual(got, test2.expected) {
			t.Errorf("ArrayDiff() = %v, want %v", got, test2.expected)
		}
	})

	test3 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"2", "4"}, {}},
		expected: []string{"1", "4"},
	}
	t.Run(test3.name, func(t *testing.T) {
		got := ArrayDiff(test3.arrays...)
		if !reflect.DeepEqual(got, test3.expected) {
			t.Errorf("ArrayDiff() = %v, want %v", got, test3.expected)
		}
	})
}

func TestArrayIntersect(t *testing.T) {
	test1 := ComparableSlice[int]{
		name:     "merge multiple arrays of integers",
		arrays:   [][]int{{1, 2}, {3, 4}, {5}},
		expected: []int{},
	}
	t.Run(test1.name, func(t *testing.T) {
		got := ArrayIntersect(test1.arrays...)
		if !reflect.DeepEqual(got, test1.expected) {
			t.Errorf("ArrayIntersect() = %v, want %v", got, test1.expected)
		}
	})
	test2 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"1", "2", "3"}},
		expected: []string{"1", "2"},
	}
	t.Run(test2.name, func(t *testing.T) {
		got := ArrayIntersect(test2.arrays...)
		if !reflect.DeepEqual(got, test2.expected) {
			t.Errorf("ArrayIntersect() = %v, want %v", got, test2.expected)
		}
	})

	test3 := ComparableSlice[string]{
		name:     "merge multiple arrays of string",
		arrays:   [][]string{{"1", "2"}, {"2", "4"}, {}},
		expected: []string{"2"},
	}
	t.Run(test3.name, func(t *testing.T) {
		got := ArrayIntersect(test3.arrays...)
		if !reflect.DeepEqual(got, test3.expected) {
			t.Errorf("ArrayIntersect() = %v, want %v", got, test3.expected)
		}
	})
}

func TestInArray(t *testing.T) {
	tests := []struct {
		name     string
		needle   any
		haystack any
		want     bool
	}{
		{"Slice - Found", 2, []int{1, 2, 3}, true},
		{"Slice - Not Found", 4, []int{1, 2, 3}, false},
		{"Array - Found", 'b', [3]rune{'a', 'b', 'c'}, true},
		{"Array - Not Found", 'd', [3]rune{'a', 'b', 'c'}, false},
		{"Map - Found Key", 2, map[string]int{"key1": 1, "key2": 2}, true},
		{"Map - Found Value", 1, map[string]int{"key1": 1, "key2": 2}, true},
		{"Map - Not Found Key", "key3", map[string]int{"key1": 1, "key2": 2}, false},
		{"Map - Not Found Value", 3, map[string]int{"key1": 1, "key2": 2}, false},
		{"Unsupported Type", 2, "string", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.needle, tt.haystack); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnique(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Test with empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Test with unique elements",
			input: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Test with duplicate elements",
			input: []int{1, 2, 2, 3, 4, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Test with nil slice",
			input: nil,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUnique(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayRand(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 5
	for i := 0; i < 10; i++ {
		result := ArrayRand(elements, n)
		fmt.Println(result)
	}
}

func TestArrayKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    map[any]any
		expected []any
	}{
		{
			name:     "Test with integers",
			input:    map[any]any{1: "one", 2: "two", 3: "three"},
			expected: []any{1, 2, 3},
		},
		{
			name:     "Test with strings",
			input:    map[any]any{"a": "apple", "b": "banana", "c": "cherry"},
			expected: []any{"a", "b", "c"},
		},
		{
			name:     "Test with empty map",
			input:    map[any]any{},
			expected: []any{},
		},
		{
			name:     "Test with nil map",
			input:    nil,
			expected: []any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayKeys(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ArrayKeys() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestArrayValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[any]any
		expected []any
	}{
		{
			name:     "Test with integers",
			input:    map[any]any{1: "one", 2: "two", 3: "three"},
			expected: []any{"one", "two", "three"},
		},
		{
			name:     "Test with empty map",
			input:    map[any]any{},
			expected: []any{},
		},
		{
			name:     "Test with strings",
			input:    map[any]any{"a": "apple", "b": "banana"},
			expected: []any{"apple", "banana"},
		},
		{
			name:     "Test with mixed types",
			input:    map[any]any{1: "one", "two": 2, 3.14: "pi"},
			expected: []any{"one", 2, "pi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ArrayValues(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("ArrayValues(%v) got %v, want %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestArrayFlip(t *testing.T) {
	arr := map[string]int{"a": 1, "b": 2, "c": 3}
	expected := map[int]string{1: "a", 2: "b", 3: "c"}
	actual := ArrayFlip(arr)
	fmt.Println(actual)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ArrayFlip(%v) = %v, want %v", arr, actual, expected)
	}
}

func TestArrayReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	actual := ArrayReverse(arr)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ArrayReverse(%v) = %v, want %v", arr, actual, expected)
	}
}

func TestArrayCountValues(t *testing.T) {
	data := []int{1, 2, 2, 3, 4, 4, 4, 5}
	expected := map[int]int{1: 1, 2: 2, 3: 1, 4: 3, 5: 1}
	actual := ArrayCountValues(data)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ArrayCountValues(%v) = %v, want %v", data, actual, expected)
	}
}

func TestShuffle(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := []int{1, 2, 3, 4, 5}
		Shuffle(arr)
		fmt.Println(arr)
	}
}
