package service

import (
	"ori/core/oriEngine"
	"ori/core/oriLog"
	"ori/internal/service/factory"
	"strings"
)

func Run(engine *oriEngine.OriEngine, services string) {
	defer engine.Wg.Done()
	f := factory.New(engine)
	if services != "" {
		replaces := strings.Replace(services, "，", ",", -1)
		servicesSlice := strings.Split(replaces, ",")
		for _, serviceName := range servicesSlice {
			service, err := f.Service(serviceName)
			if err != nil {
				oriLog.Error("服务不存在:%s", serviceName)
				continue
			}
			engine.Wg.Add(1)
			go service.Run(engine)
		}
	}
}
