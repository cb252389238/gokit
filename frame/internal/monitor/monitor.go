package monitor

import (
	"fmt"
	"frame/internal/svc"
	"frame/resource"
	"frame/util/common"
	"runtime"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	webHookHttpClient = resty.New() //请求实力
	monitorFluctuate  = map[string]interface{}{}
	monitorLock       sync.Mutex
)

const (
	MAX_CPU_PERCENT     = 60   //cpu最大使用率
	MAX_MEM_PERCENT     = 60   //内存最大使用率
	MAX_DISK_PERCENT    = 70   //硬盘最大使用率
	MAX_GOROUTINE_NUM   = 3000 //协程最大数量
	PERCENT_FLUCTUATE   = 5    //波动百分比 超过阈值则进行通知，在范围内不通知
	GOROUTINE_FLUCTUATE = 500  //协程数量波动正常范围阈值
)

func MonitorService(ctx *svc.ServiceContext) {
	defer ctx.Wg.Done()
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ctx.Context.Done():
			resource.LogInfo("监控协程退出")
			return
		case <-ticker.C:
			cpuPercent := common.GetCpuPercent()   //cpu使用率
			memPercent := common.GetMemPercent()   //内存使用率
			diskPercent := common.GetDiskPercent() //硬盘使用率
			goroutineNum := runtime.NumGoroutine() //开启协程数量
			if fluctuateMargin("cpuPercent", cpuPercent) {
				fmt.Sprintf("CPU使用率:%.2f", cpuPercent)
			}
			if fluctuateMargin("memPercent", memPercent) {
				fmt.Sprintf("内存使用率:%.2f", memPercent)
			}
			if fluctuateMargin("diskPercent", diskPercent) {
				fmt.Sprintf("硬盘使用率:%.2f", diskPercent)
			}
			if fluctuateMargin("goroutineNum", goroutineNum) {
				fmt.Sprintf("协程数量:%d", goroutineNum)
			}
		}
	}
}

// 检测波动是否超出阈值并通知
func fluctuateMargin(t string, v interface{}) bool {
	monitorLock.Lock()
	defer monitorLock.Unlock()
	b := false
	switch t {
	case "cpuPercent":
		if _, ok := monitorFluctuate[t]; ok {
			cpuPercent := v.(float64)
			b = cpuPercent > MAX_CPU_PERCENT && (cpuPercent < monitorFluctuate[t+"lower"].(float64) || cpuPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
			}
		} else {
			b = v.(float64) > MAX_CPU_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "memPercent":
		if _, ok := monitorFluctuate[t]; ok {
			memPercent := v.(float64)
			b = memPercent > MAX_MEM_PERCENT && (memPercent < monitorFluctuate[t+"lower"].(float64) || memPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
			}
		} else {
			b = v.(float64) > MAX_MEM_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "diskPercent":
		if _, ok := monitorFluctuate[t]; ok {
			diskPercent := v.(float64)
			b = diskPercent > MAX_DISK_PERCENT && (diskPercent < monitorFluctuate[t+"lower"].(float64) || diskPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
			}
		} else {
			b = v.(float64) > MAX_DISK_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - PERCENT_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + PERCENT_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "goroutineNum":
		if _, ok := monitorFluctuate[t]; ok {
			goroutineNum := v.(int)
			b = goroutineNum > MAX_GOROUTINE_NUM && (goroutineNum < monitorFluctuate[t+"lower"].(int) || goroutineNum > monitorFluctuate[t+"upper"].(int))
			if b {
				monitorFluctuate[t+"lower"] = v.(int) - GOROUTINE_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(int) + GOROUTINE_FLUCTUATE
			}
		} else {
			b = v.(int) > MAX_GOROUTINE_NUM
			monitorFluctuate[t+"lower"] = v.(int) - GOROUTINE_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(int) + GOROUTINE_FLUCTUATE
		}
		monitorFluctuate[t] = v
	}
	return b
}