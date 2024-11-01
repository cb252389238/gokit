package redisKey

import (
	"fmt"
)

var (
	LockUserRegister = func(mobile, regionCode string) string {
		return fmt.Sprintf("userRegisterLock:%v:%v", regionCode, mobile)
	}
)
