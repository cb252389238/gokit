package oriMonitor

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
	"ori/core/oriConfig"
	"ori/core/oriLog"
	oriTools2 "ori/core/oriTools/concurrence"
	"ori/core/oriTools/easy"
	"ori/internal/engine"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	monitorFluctuate = map[string]interface{}{}
	monitorLock      sync.Mutex
	lastMsgTime      int64
)

func SendNotice(message string) {
	webHookMsgChan <- webHookMsgTextData{
		Content: message,
	}
}

func Monitor(ctx *engine.OriEngine) {
	defer ctx.Wg.Done()
	ticker := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-ctx.Context.Done():
			oriLog.LogInfo("监控服务退出")
			return
		case data := <-webHookMsgChan:
			if time.Now().Unix()-lastMsgTime <= 20 {
				continue
			}
			ip := easy.GetLoaclIp()
			allConfig := oriConfig.GetHotConf()
			data.Content = fmt.Sprintf("【%s】[ip:%s]\r\n%s", allConfig.APP, ip, data.Content)
			if strings.ToUpper(allConfig.ENV) == "PRO" {
				//生产环境
				ctx.WebHook.SendTextMessage(data.Content, dingtalk.WithAtAll())
			} else {
				oriLog.LogInfo("%+v", data)
			}
			lastMsgTime = time.Now().Unix()
		case <-ticker.C:
			cpuPercent := GetCpuPercent()          //cpu使用率
			memPercent := GetMemPercent()          //内存使用率
			diskPercent := GetDiskPercent()        //硬盘使用率
			goroutineNum := runtime.NumGoroutine() //开启协程数量
			if fluctuateMargin("cpuPercent", cpuPercent) {
				SendNotice(fmt.Sprintf("CPU使用率:%.2f", cpuPercent))
			}
			if fluctuateMargin("memPercent", memPercent) {
				SendNotice(fmt.Sprintf("内存使用率:%.2f", memPercent))
			}
			if fluctuateMargin("diskPercent", diskPercent) {
				SendNotice(fmt.Sprintf("硬盘使用率:%.2f", diskPercent))
			}
			if fluctuateMargin("goroutineNum", goroutineNum) {
				SendNotice(fmt.Sprintf("协程数量:%d", goroutineNum))
			}
			//接口并发数量监控
			oriTools2.ConcurrencyNum.Range(func(key, value interface{}) bool {
				routerName := key.(string)
				num := value.(int)
				if fluctuateMargin(routerName, value) {
					if routerName == "all" {
						SendNotice(fmt.Sprintf("服务器总并发数:%d", num))
					}
					if routerName != "all" {
						SendNotice(fmt.Sprintf("接口:[%s]并发数:%d", routerName, num))
					}
				}
				return true
			})
		}
	}
}

// 检测波动是否超出阈值并通知
func fluctuateMargin(t string, v interface{}) bool {
	monitorLock.Lock()
	defer monitorLock.Unlock()
	monitorConfig := oriConfig.GetHotConf().Monitor
	b := false
	switch t {
	case "cpuPercent":
		if _, ok := monitorFluctuate[t]; ok {
			cpuPercent := v.(float64)
			b = cpuPercent > monitorConfig.MAX_CPU_PERCENT && (cpuPercent < monitorFluctuate[t+"lower"].(float64) || cpuPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.CPU_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.CPU_FLUCTUATE
			}
		} else {
			b = v.(float64) > monitorConfig.MAX_CPU_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.CPU_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.CPU_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "memPercent":
		if _, ok := monitorFluctuate[t]; ok {
			memPercent := v.(float64)
			b = memPercent > monitorConfig.MAX_MEM_PERCENT && (memPercent < monitorFluctuate[t+"lower"].(float64) || memPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.MEM_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.MEM_FLUCTUATE
			}
		} else {
			b = v.(float64) > monitorConfig.MAX_MEM_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.MEM_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.MEM_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "diskPercent":
		if _, ok := monitorFluctuate[t]; ok {
			diskPercent := v.(float64)
			b = diskPercent > monitorConfig.MAX_DISK_PERCENT && (diskPercent < monitorFluctuate[t+"lower"].(float64) || diskPercent > monitorFluctuate[t+"upper"].(float64))
			if b {
				monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.DISK_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.DISK_FLUCTUATE
			}
		} else {
			b = v.(float64) > monitorConfig.MAX_DISK_PERCENT
			monitorFluctuate[t+"lower"] = v.(float64) - monitorConfig.DISK_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(float64) + monitorConfig.DISK_FLUCTUATE
		}
		monitorFluctuate[t] = v
	case "goroutineNum":
		if _, ok := monitorFluctuate[t]; ok {
			goroutineNum := v.(int)
			b = goroutineNum > monitorConfig.MAX_GOROUTINE_NUM && (goroutineNum < monitorFluctuate[t+"lower"].(int) || goroutineNum > monitorFluctuate[t+"upper"].(int))
			if b {
				monitorFluctuate[t+"lower"] = v.(int) - monitorConfig.GOROUTINE_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(int) + monitorConfig.GOROUTINE_FLUCTUATE
			}
		} else {
			b = v.(int) > monitorConfig.MAX_GOROUTINE_NUM
			monitorFluctuate[t+"lower"] = v.(int) - monitorConfig.GOROUTINE_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(int) + monitorConfig.GOROUTINE_FLUCTUATE
		}
		monitorFluctuate[t] = v
	default:
		if _, ok := monitorFluctuate[t]; ok {
			concurrencyNum := v.(int)
			b = concurrencyNum > monitorConfig.MAX_CONCURRENCY_NUM && (concurrencyNum < monitorFluctuate[t+"lower"].(int) || concurrencyNum > monitorFluctuate[t+"upper"].(int))
			if b {
				monitorFluctuate[t+"lower"] = v.(int) - monitorConfig.CONCURRENCY_FLUCTUATE
				monitorFluctuate[t+"upper"] = v.(int) + monitorConfig.CONCURRENCY_FLUCTUATE
			}
		} else {
			b = v.(int) > monitorConfig.MAX_CONCURRENCY_NUM
			monitorFluctuate[t+"lower"] = v.(int) - monitorConfig.CONCURRENCY_FLUCTUATE
			monitorFluctuate[t+"upper"] = v.(int) + monitorConfig.CONCURRENCY_FLUCTUATE
		}
		monitorFluctuate[t] = v
	}
	return b
}
