package core

import (
	"log"
	"sync"
	"train-mq/models"
)

// MainMessageQueue 可并发消息队列结构体
type MainMessageQueue struct {
	topics      sync.Map // 存储主题和对应的无锁队列
	subscribers sync.Map // 存储主题和订阅者信息
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

// RegisterSubscriber 注册订阅者
func (mmq *MainMessageQueue) RegisterSubscriber(topic string, subscriber models.Subscriber) {
	log.Printf("Attempting to register subscriber: %v to topic: %s", subscriber, topic)
	hasSubscriber := mmq.HasSubscriber(topic, subscriber.SubId)
	if hasSubscriber {
		// 如果已经订阅过则忽略
		log.Printf("Subscriber: %v is already registered, ignoring", subscriber)
		return
	}
	// 如果没有订阅过则添加订阅者信息
	subscribers, _ := mmq.subscribers.Load(topic)
	if subscribers == nil {
		subscribers = make([]models.Subscriber, 0)
	}
	existingSubscribers := subscribers.([]models.Subscriber)
	if existingSubscribers == nil {
		existingSubscribers = make([]models.Subscriber, 0)
	}
	existingSubscribers = append(existingSubscribers, subscriber)
	mmq.subscribers.Store(topic, existingSubscribers)
	log.Printf("Successfully registered subscriber: %v to topic: %s", subscriber, topic)

}

// HasSubscriber 判断主题是否有指定订阅者
func (mmq *MainMessageQueue) HasSubscriber(topic string, subId string) bool {
	subscribers, _ := mmq.subscribers.Load(topic)
	if subscribers == nil {
		return false
	}
	existingSubscribers := subscribers.([]models.Subscriber)
	for _, existingSubscriber := range existingSubscribers {
		if existingSubscriber.SubId == subId {
			return true
		}
	}
	return false
}

// GetSubscribers 获取指定主题的所有订阅者
func (mmq *MainMessageQueue) GetSubscribers(topic string) []models.Subscriber {
	subscribers, _ := mmq.subscribers.Load(topic)
	return subscribers.([]models.Subscriber)
}
