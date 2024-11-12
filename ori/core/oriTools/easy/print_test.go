package easy

import (
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	Echo("123", "456")
}

func TestVar_dump(t *testing.T) {
	Var_dump(map[string]int64{"aa": 1})
}

func TestExit(t *testing.T) {
	Exit(1)
}

func TestDie(t *testing.T) {
	Die(1)
}

func TestGetenv(t *testing.T) {
	fmt.Println(Getenv("GOPATH"))
}

func TestPutenv(t *testing.T) {
	Putenv("test_env=1")
}

func TestVersion_compare(t *testing.T) {
	fmt.Println(Version_compare("1.0.0.1", "1.0.0.2", "<"))
}
