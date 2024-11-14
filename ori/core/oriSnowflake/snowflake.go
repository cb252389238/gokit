package oriSnowflake

import (
	snowflake2 "ori/core/oriTools/snowflake"
	"sync"
)

var (
	once      sync.Once
	snowflake *snowflake2.Node
	err       error
)

func New(node int64) error {
	once.Do(func() {
		snowflake, err = snowflake2.NewNode(node)
	})
	return err
}

// 获取雪花id
func GetSnowId() string {
	return snowflake.Generate().String()
}
