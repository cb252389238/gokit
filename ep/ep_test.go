package ep

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	fmt.Println(InArray(1, []int{1, 2, 3, 4}))
	fmt.Println(InArray(1, []int{2, 3, 4, 1}))
	fmt.Println(InArray(1, []int{2, 3, 4}))
}
