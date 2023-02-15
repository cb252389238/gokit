package factory

import (
	"errors"
	"fmt"
	"ori/typedef"
)

var (
	serviceMap map[string]Service //map存储所有服务
)

// 需要实现的方法
type Service interface {
	Example(req *typedef.Request, resp *typedef.Response)
}

// 工厂类型
type Factory struct {
}

// 获取对应的服务
func (f *Factory) Service(serviceName string) (Service, error) {
	if v, ok := serviceMap[serviceName]; ok {
		return v, nil
	}
	return nil, errors.New(fmt.Sprintf("服务不存在 serviceName:%s", serviceName))
}

func New() *Factory {
	serviceMap = map[string]Service{
		"example": new(example),
	}
	return new(Factory)
}
