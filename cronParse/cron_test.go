package cronParse

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	expr, err := Parse("*/10 * * * * * * ")
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	next := expr.Next(now)
	fmt.Println(now.Unix())
	fmt.Println(next.Unix())
}
