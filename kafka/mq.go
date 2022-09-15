package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 生产者
func producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.NoResponse              //不需要确认
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner //轮询
	config.Producer.Return.Successes = false                      //同步模式失败和成功一般都开启
	config.Producer.Return.Errors = true                          //返回错误
	producer, err := sarama.NewAsyncProducer([]string{"47.96.156.121:9092"}, config)
	//defer producer.Close()
	defer producer.AsyncClose()
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic:     "topic_chat_room", //主题
		Key:       sarama.StringEncoder("key"),
		Value:     sarama.StringEncoder("golang kafka 测试消息"),
		Offset:    0,
		Partition: 0,
	}
	//同步发送消息
	//分区id 偏移量 错误
	//pid, offset, err := producer.SendMessage(msg)//同步模式
	//fmt.Println(pid, offset, err)
	//异步发送消息
	producer.Input() <- msg
	select {
	case success := <-producer.Successes():
		fmt.Println(success)
	case err := <-producer.Errors():
		fmt.Println(err)
	}
}

// 消费者
func consumer() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
