package php

import (
	"fmt"
	"testing"
)

func TestAbs(t *testing.T) {
	a := 123
	fmt.Println(Abs(a))
	b := -111
	fmt.Println(Abs(b))
}

func TestCeil(t *testing.T) {
	fmt.Println(Ceil(3.1415))
}

func TestFloor(t *testing.T) {
	fmt.Println(Floor(3.1415))
}

func TestMax(t *testing.T) {
	fmt.Println(Max(1, 2, 3, 4, 5))
	fmt.Println(Max(1.1, 5.5, 3.2, 9.8, 100.0))
}

func TestMin(t *testing.T) {
	fmt.Println(Min(1, 2, 3, 4, 5))
	fmt.Println(Min(1.1, 5.5, 3.2, 9.8, 100.0))
}

func TestMt_rand(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(Mt_rand(5, 10))
	}
}
