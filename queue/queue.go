package queue

import (
	"sync"
	"trainMQ/models"
)

// MessageQueue 消息队列
type MessageQueue struct {
	// 主题 -> 通道
	messages map[string]chan models.Message
	// 读写锁
	mux sync.RWMutex
}

// NewMessageQueue 初始化消息队列
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		messages: make(map[string]chan models.Message),
	}
}

// Publish 发布消息到主题
func (mq *MessageQueue) Publish(msg models.Message) {
	mq.mux.Lock()
	defer mq.mux.Unlock()
	// 如果主题不存在则会被创建
	if _, exists := mq.messages[msg.Topic]; !exists {
		mq.messages[msg.Topic] = make(chan models.Message, 100)
	}
	mq.messages[msg.Topic] <- msg
}

// Consume 消费某个主题的消息
func (mq *MessageQueue) Consume(topic string) (models.Message, bool) {
	mq.mux.RLock()
	defer mq.mux.RUnlock()

	if ch, exists := mq.messages[topic]; exists {
		select {
		case msg := <-ch:
			return msg, true
		default:
			return models.Message{}, false
		}
	}
	// 如果没有消息则返回 false
	return models.Message{}, false
}
