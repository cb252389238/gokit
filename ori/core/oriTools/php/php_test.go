package php

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	a := ""
	fmt.Println(Empty(a))
	b := []int{1}
	fmt.Println(Empty(b))
	c := []int{}
	fmt.Println(Empty(c))
	d := map[string]int{}
	fmt.Println(Empty(d))
	e := map[string]any{"aa": 1}
	fmt.Println(Empty(e))
}

func TestIs_numeric(t *testing.T) {
	a := "123"
	fmt.Println(Is_numeric(a))
	b := "abc"
	fmt.Println(Is_numeric(b))
	c := 100
	fmt.Println(Is_numeric(c))
}
