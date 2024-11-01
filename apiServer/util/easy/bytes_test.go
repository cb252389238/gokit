package easy

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestBinaryEncode(t *testing.T) {
	str := "123456789"
	encode, err := BinaryEncode(str)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("encode:", encode)
	decode, err := BinaryDecode(encode)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(decode)
}
