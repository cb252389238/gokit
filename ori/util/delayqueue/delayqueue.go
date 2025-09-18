package core_delayqueue

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

const (
	DELAYQUEUE = "DELAYQUEUE"
)

type Task struct {
	Id      string    // 任务唯一标识
	Payload any       // 任务数据
	Delay   time.Time // 延时时间
}

// 延时队列
type DelayQueue struct {
	client *redis.Client
	name   string // 队列名称
	C      chan Task
}

// 创建新的延时队列
func NewDelayQueue(client *redis.Client, name string) *DelayQueue {
	dq := &DelayQueue{
		client: client,
		name:   fmt.Sprintf("%s:%s", DELAYQUEUE, name),
		C:      make(chan Task, 1000),
	}
	go dq.consumer()
	return dq
}

// 添加延时任务
func (dq *DelayQueue) AddTask(task Task) error {
	if task.Id == "" {
		return errors.New("task ID cannot be empty")
	}
	_, err := dq.client.ZAdd(dq.name, redis.Z{
		Score:  float64(task.Delay.Unix()),
		Member: task.Id,
	}).Result()
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(task)
	if err != nil {
		return err
	}
	_, err = dq.client.Set(dq.getTaskKey(task.Id), string(marshal), time.Second*time.Duration(task.Delay.Unix()-time.Now().Unix()+60)).Result()
	return err
}

// 移除延时任务
func (dq *DelayQueue) RemoveTask(id string) error {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	_, err := dq.client.ZRem(dq.name, id).Result()
	if err != nil {
		return err
	}
	_, err = dq.client.Del(dq.getTaskKey(id)).Result()
	if err != nil {
		return err
	}
	return nil
}

// 获取到期的任务
func (dq *DelayQueue) Poll() ([]Task, error) {
	now := time.Now().Unix()
	// 获取所有score小于等于当前时间戳的任务
	taskIds, err := dq.client.ZRangeByScore(dq.name, redis.ZRangeBy{
		Min: "-inf",
		Max: cast.ToString(now),
	}).Result()
	if err != nil {
		return nil, err
	}
	if len(taskIds) == 0 {
		return nil, nil
	}
	var tasks []Task
	for _, taskId := range taskIds {
		success, unlock, err := dq.lock(dq.getLockKey(taskId), time.Second)
		if !success {
			continue
		}
		if err != nil {
			continue
		}
		result, _ := dq.client.Get(dq.getTaskKey(taskId)).Result()
		t := Task{}
		if result != "" {
			_ = json.Unmarshal([]byte(result), &t)
		}

		delNum, _ := dq.client.ZRem(dq.name, taskId).Result()
		if delNum == 0 {
			unlock()
			continue
		}
		tasks = append(tasks, t)
		unlock()
	}
	return tasks, nil
}

// 启动消费者
func (dq *DelayQueue) consumer() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			tasks, err := dq.Poll()
			if err != nil {
				continue
			}
			if len(tasks) > 0 {
				for _, task := range tasks {
					dq.C <- task
				}
			}
		}
	}
}

// 获取任务存储的key
func (dq *DelayQueue) getTaskKey(id string) string {
	return dq.name + ":store:" + id
}

func (dq *DelayQueue) getLockKey(id string) string {
	return dq.name + ":lock:" + id
}

func (dq *DelayQueue) lock(key string, lockTime time.Duration) (success bool, unlock func(), err error) {
	value := time.Now().UnixMicro()
	val := strconv.FormatInt(value, 10)
	stm, err := dq.client.SetNX(key, val, lockTime).Result()
	if err != nil {
		return false, nil, err
	}
	if stm == false {
		return false, nil, nil
	}
	return true, func() {
		dq.client.Del(key)
	}, nil
}
