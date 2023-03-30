package factory

import (
	"errors"
	"fmt"
	"ori/core/oriEngine"
	"ori/internal/service/task"
)

var (
	serviceMap map[string]Service //map存储所有服务
)

// 需要实现的方法
type Service interface {
	Run(engine *oriEngine.OriEngine)
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

func New(engine *oriEngine.OriEngine) *Factory {
	serviceMap = map[string]Service{
		"example": new(task.Example),
	}
	return new(Factory)
}
