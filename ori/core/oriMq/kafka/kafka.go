package kafka

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"ori/core/oriLog"
	"time"
)

type Kafka struct {
	*kafka
}

type kafka struct {
	asyncProducer sarama.AsyncProducer //异步生产者
	syncProducer  sarama.SyncProducer  //同步生产者
	consumer      sarama.ConsumerGroup
	errChan       chan error //错误通道
	context       context.Context
	topic         string
	addr          []string
	group         string
	*consume
}

type KafkaConf struct {
	Addr       []string        //实例地址
	Group      string          //分组名
	Topic      string          //主题
	Context    context.Context //上下文
	InstanceId string          //静态成员id
}

type ConsumerMessage struct {
	Message   string
	Timestamp time.Time
	Key       string
	Topic     string
	Partition int32
	Offset    int64
	Session   *sarama.ConsumerMessage
}

type consume struct {
	C       chan ConsumerMessage //通道
	session sarama.ConsumerGroupSession
	quit    chan int
	Close   chan int
}

// 方法在开始时运行
func (c *consume) Setup(sarama.ConsumerGroupSession) error {
	fmt.Println("Setup")
	return nil
}

// 方法在结束时运行
func (c *consume) Cleanup(sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup")
	return nil
}

// 消费循环
func (c *consume) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	c.session = session
	fmt.Println("start")
	defer fmt.Println("end")
	for {
		select {
		case <-c.Close:
			log.Printf("kafka收到退出信号")
			return nil
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			c.C <- ConsumerMessage{
				Message:   string(message.Value),
				Timestamp: message.Timestamp,
				Key:       string(message.Key),
				Topic:     message.Topic,
				Partition: message.Partition,
				Offset:    message.Offset,
				Session:   message,
			}
		}
	}
}

func New(conf KafkaConf) (*Kafka, error) {
	k := &Kafka{
		kafka: &kafka{
			errChan: make(chan error, 10),
			consume: &consume{
				C:     make(chan ConsumerMessage, 100),
				quit:  make(chan int),
				Close: make(chan int),
			},
		},
	}
	//异步生产者
	//{
	//	config := sarama.NewConfig()
	//	config.Producer.RequiredAcks = sarama.NoResponse
	//	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //轮训
	//	config.Producer.Return.Successes = false                      // 成功交付的消息将在success_channel返回
	//	config.Producer.Return.Errors = true
	//	client, err := sarama.NewAsyncProducer(conf.Addr, config)
	//	if err != nil {
	//		return nil, err
	//	}
	//	k.asyncProducer = client
	//}
	//同步生产者

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //轮训
	config.Producer.Return.Successes = true                       // 成功交付的消息将在success_channel返回
	config.Producer.Return.Errors = true
	client, err := sarama.NewSyncProducer(conf.Addr, config)
	if err != nil {
		return nil, err
	}
	k.syncProducer = client

	//消费者
	{
		config := sarama.NewConfig()
		//sarama.BalanceStrategyRoundRobin
		//sarama.BalanceStrategyRange
		//sarama.BalanceStrategySticky
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
			sarama.NewBalanceStrategySticky(),
		}
		config.Consumer.Return.Errors = true
		config.Consumer.Offsets.AutoCommit.Enable = true
		config.Consumer.Offsets.AutoCommit.Interval = time.Millisecond * 100
		config.Consumer.Offsets.Initial = sarama.OffsetNewest
		if conf.InstanceId != "" {
			config.Consumer.Group.InstanceId = conf.InstanceId
		}
		newClient, err := sarama.NewClient(conf.Addr, config)
		if err != nil {
			return nil, err
		}
		client, err := sarama.NewConsumerGroupFromClient(conf.Group, newClient)
		if err != nil {
			return nil, err
		}
		k.consumer = client
	}
	k.context = conf.Context
	k.addr = conf.Addr
	k.topic = conf.Topic
	k.group = conf.Group
	go k.run() //执行常驻
	return k, nil
}

// 执行程序
func (k *kafka) run() {
	go func() {
		for {
			select {
			case <-k.context.Done():
				return
			case <-time.After(time.Second * 3):
				if err := k.consumer.Consume(k.context, []string{k.topic}, k.consume); err != nil {
					oriLog.Error("kafka再平衡计算错误 err:%+v,topic:%s", err, k.topic)
				} else {
					oriLog.Info("kafka再平衡计算 topic:%+v", k.topic)
				}
			}
		}
	}()
	for {
		select {
		case <-k.context.Done():
			close(k.quit)
			close(k.Close)
			k.consumer.Close()
			oriLog.Info("kafka退出完成 topic:%s", k.topic)
			return
		}
	}
}

func (k *kafka) Ack(message *sarama.ConsumerMessage) {
	k.session.MarkMessage(message, "")
	//k.session.Commit()
}

func (k *kafka) Pull() (ConsumerMessage, error) {
	select {
	case <-time.After(time.Millisecond * 10): //超时处理 十毫秒没有消息返回
		return ConsumerMessage{}, errors.New("not found new message")
	case message := <-k.C:
		return message, nil
	}
}

// 同步写入消息
func (k *kafka) Push(msg []string) error {
	if len(msg) <= 0 {
		return errors.New("msg is empty")
	}
	data := make([]*sarama.ProducerMessage, 0, len(msg))
	for _, v := range msg {
		data = append(data, &sarama.ProducerMessage{
			Topic: k.topic,
			Value: sarama.StringEncoder(v),
		})
	}
	err := k.syncProducer.SendMessages(data)
	return err
}
