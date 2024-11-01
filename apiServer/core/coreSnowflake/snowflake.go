package coreSnowflake

import (
	snowflake2 "apiServer/core/coreTools/snowflake"
	"sync"
)

var (
	once      sync.Once
	snowflake *snowflake2.Node
	err       error
)

func New(node int64) (*snowflake2.Node, error) {
	once.Do(func() {
		snowflake, err = snowflake2.NewNode(node)
	})
	return snowflake, err
}

func GetSnowId() string {
	return snowflake.Generate().String()
}
