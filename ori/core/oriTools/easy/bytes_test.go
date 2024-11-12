package easy

import (
	"testing"
)

func TestBinary(t *testing.T) {
	msg := "hello world"
	binary, _ := BinaryEncode(msg)
	t.Log(binary)
	msg2, _ := BinaryDecode(binary)
	t.Log(msg2)
}
