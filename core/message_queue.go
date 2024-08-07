package core

import (
	"log"
	"sync"
	"train-mq/models"
)

// MainContext 消息队列上下文
type MainContext struct {
	topicsQueue      sync.Map // 存储主题和对应的无锁队列
	topicConsumerMap sync.Map // 存储主题和消费者信息
	cidMap           sync.Map // 存储消费者信息，cid -> consumer
}

// NewMainContext 创建一个新的上下文
func NewMainContext() *MainContext {
	return &MainContext{}
}

// GetQueue 获取指定主题的队列
func (mmq *MainContext) GetQueue(topic string) (*models.LockFreeQueue, bool) {
	queue, ok := mmq.topicsQueue.Load(topic)
	if !ok {
		return nil, false
	}
	return queue.(*models.LockFreeQueue), true
}

// Publish 发布消息到主题
func (mmq *MainContext) Publish(message models.Message) {
	queue, ok := mmq.GetQueue(message.Topic)
	if !ok {
		// 如果主题不存在则会被创建
		queue = models.NewLockFreeQueue()
		mmq.topicsQueue.Store(message.Topic, queue)
	}
	queue.Enqueue(message)
}

// Consume 消费某个主题的消息
func (mmq *MainContext) Consume(topic string, offset int) (models.Message, bool) {
	queue, ok := mmq.GetQueue(topic)
	if !ok {
		return models.Message{}, false
	}
	return queue.PeekAt(offset)
}

// RegisterConsumer 注册消费者
func (mmq *MainContext) RegisterConsumer(topic string, cid string) {
	log.Printf("Attempting to register cid: %s to topic: %s", cid, topic)
	newConsumer := models.NewConsumer(cid)
	newConsumer.SetOffset(topic, 0)

	// 如果订阅过则忽略，没订阅过则存储消费者信息
	hasConsumer := mmq.HasConsumer(topic, cid)
	if hasConsumer {
		log.Printf("Cid: %v is already registered, ignoring", cid)
		return
	} else {
		// 添加主题的消费者信息
		consumers, _ := mmq.topicConsumerMap.LoadOrStore(topic, make([]models.Consumer, 0, 1))
		consumerList := consumers.([]models.Consumer)
		consumerList = append(consumerList, *newConsumer)
		mmq.topicConsumerMap.Store(topic, consumerList)
		// 添加系统的消费者
		sysConsumer, loaded := mmq.cidMap.LoadOrStore(cid, newConsumer)
		if loaded {
			// 如果消费者已存在，添加新主题
			existingConsumer := sysConsumer.(*models.Consumer)
			existingConsumer.TopicOffsetMap.Store(topic, 0)
			mmq.cidMap.Store(cid, existingConsumer)
		}

		log.Printf("Successfully registered consumer. cid: %v to topic: %s", cid, topic)
	}

}

// HasConsumer 判断主题是否有指定消费者
func (mmq *MainContext) HasConsumer(topic string, cid string) bool {
	consumers, _ := mmq.topicConsumerMap.Load(topic)
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

// GetConsumersByTopic 获取指定主题的所有消费者
func (mmq *MainContext) GetConsumersByTopic(topic string) []models.Consumer {
	consumers, _ := mmq.topicConsumerMap.Load(topic)
	return consumers.([]models.Consumer)
}

// GetConsumerByCid 获取指定消费者信息
func (mmq *MainContext) GetConsumerByCid(cid string) *models.Consumer {
	consumer, _ := mmq.cidMap.Load(cid)
	return consumer.(*models.Consumer)
}
