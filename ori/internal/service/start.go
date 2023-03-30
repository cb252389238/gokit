package service

import (
	log "github.com/sirupsen/logrus"
	"ori/core/oriEngine"
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
				log.Error("打印日志错误")
			}
			engine.Wg.Add(1)
			go service.Run(engine)
		}
	}
}
