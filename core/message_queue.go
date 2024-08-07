package core

import (
	"log"
	"sync"
	"train-mq/models"
)

// MainMessageQueue 消息队列
type MainMessageQueue struct {
	topicsQueue    sync.Map // 存储主题和对应的无锁队列
	topicConsumers sync.Map // 存储主题和消费者信息
	consumers      sync.Map // 存储消费者信息，cid -> consumer
}

// NewMainMessageQueue 创建一个新的消息队列
func NewMainMessageQueue() *MainMessageQueue {
	return &MainMessageQueue{}
}

// GetQueue 获取指定主题的队列
func (mmq *MainMessageQueue) GetQueue(topic string) (*models.LockFreeQueue, bool) {
	queue, ok := mmq.topicsQueue.Load(topic)
	if !ok {
		return nil, false
	}
	return queue.(*models.LockFreeQueue), true
}

// Publish 发布消息到主题
func (mmq *MainMessageQueue) Publish(message models.Message) {
	queue, ok := mmq.GetQueue(message.Topic)
	if !ok {
		// 如果主题不存在则会被创建
		queue = models.NewLockFreeQueue()
		mmq.topicsQueue.Store(message.Topic, queue)
	}
	queue.Enqueue(message)
}

// Consume 消费某个主题的消息
func (mmq *MainMessageQueue) Consume(topic string) (models.Message, bool) {
	queue, ok := mmq.GetQueue(topic)
	if !ok {
		return models.Message{}, false
	}
	return queue.Dequeue()
}

// RegisterConsumer 注册消费者
func (mmq *MainMessageQueue) RegisterConsumer(topic string, consumer models.Consumer) {
	log.Printf("Attempting to register consumer: %v to topic: %s", consumer, topic)
	hasConsumer := mmq.HasConsumer(topic, consumer.Cid)
	// 如果订阅过则忽略，没订阅过则存储消费者信息
	if hasConsumer {
		log.Printf("Consumer: %v is already registered, ignoring", consumer)
		return
	} else {
		// 添加主题的消费者信息
		consumers, _ := mmq.topicConsumers.LoadOrStore(topic, make([]models.Consumer, 0, 1))
		consumerList := consumers.([]models.Consumer)
		consumerList = append(consumerList, consumer)
		mmq.topicConsumers.Store(topic, consumerList)
		// 添加系统的消费者
		sysConsumer, loaded := mmq.consumers.LoadOrStore(consumer.Cid, consumer)
		if loaded {
			// 如果消费者已存在，添加新主题
			existingConsumer := sysConsumer.(models.Consumer)
			existingConsumer.Topics = append(existingConsumer.Topics, topic)
			mmq.consumers.Store(consumer.Cid, existingConsumer)
		}

		log.Printf("Successfully registered consumer: %v to topic: %s", consumer, topic)
	}

}

// HasConsumer 判断主题是否有指定消费者
func (mmq *MainMessageQueue) HasConsumer(topic string, cid string) bool {
	consumers, _ := mmq.topicConsumers.Load(topic)
	if consumers == nil {
		return false
	}
	existingConsumers := consumers.([]models.Consumer)
	for _, existingConsumer := range existingConsumers {
		if existingConsumer.Cid == cid {
			return true
		}
	}
	return false
}

// GetConsumers 获取指定主题的所有消费者
func (mmq *MainMessageQueue) GetConsumers(topic string) []models.Consumer {
	consumers, _ := mmq.topicConsumers.Load(topic)
	return consumers.([]models.Consumer)
}
