package chscht

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	dicter, err := New()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dicter.Traditional("中华人民共和国"))
}
