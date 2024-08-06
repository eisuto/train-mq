package core

import (
	"log"
	"sync"
	"train-mq/models"
)

// MainMessageQueue 可并发消息队列结构体
type MainMessageQueue struct {
	topics    sync.Map // 存储主题和对应的无锁队列
	consumers sync.Map // 存储主题和消费者信息
}

// NewMainMessageQueue 创建一个新的可并发消息队列
func NewMainMessageQueue() *MainMessageQueue {
	return &MainMessageQueue{}
}

// GetQueue 获取指定主题的队列
func (mmq *MainMessageQueue) GetQueue(topic string) (*models.LockFreeQueue, bool) {
	queue, ok := mmq.topics.Load(topic)
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
		mmq.topics.Store(message.Topic, queue)
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
	if hasConsumer {
		// 如果已经订阅过则忽略
		log.Printf("Consumer: %v is already registered, ignoring", consumer)
		return
	}
	// 如果没有订阅过则添加消费者信息
	consumers, _ := mmq.consumers.Load(topic)
	if consumers == nil {
		consumers = make([]models.Consumer, 0)
	}
	existingConsumers := consumers.([]models.Consumer)
	if existingConsumers == nil {
		existingConsumers = make([]models.Consumer, 0)
	}
	existingConsumers = append(existingConsumers, consumer)
	mmq.consumers.Store(topic, existingConsumers)
	log.Printf("Successfully registered consumer: %v to topic: %s", consumer, topic)

}

// HasConsumer 判断主题是否有指定消费者
func (mmq *MainMessageQueue) HasConsumer(topic string, cid string) bool {
	consumers, _ := mmq.consumers.Load(topic)
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
	consumers, _ := mmq.consumers.Load(topic)
	return consumers.([]models.Consumer)
}
