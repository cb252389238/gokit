package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	jwtSecret := "123456"
	expireTime := time.Now().Add(time.Hour * 3).Unix()
	c := Claims{
		UserId: "1001",
	}
	token, err := Encode(c, []byte(jwtSecret), expireTime)
	fmt.Println(token, err)

	user, err := Decode(token, []byte(jwtSecret))
	fmt.Println(user, err)
}

func BenchmarkToken(b *testing.B) {
	jwtSecret := "123456"
	expireTime := time.Now().Add(time.Hour * 3).Unix()
	c := Claims{
		UserId: "1001",
	}
	for i := 0; i < b.N; i++ {
		token, _ := Encode(c, []byte(jwtSecret), expireTime)
		Decode(token, []byte(jwtSecret))
	}
}
