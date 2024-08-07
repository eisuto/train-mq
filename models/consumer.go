package models

import "sync"

// Consumer 消费者结构体
type Consumer struct {
	Cid            string    `json:"cid"`            // 消费者ID
	TopicOffsetMap *sync.Map `json:"topicOffsetMap"` // 消费者的 主题 -> 偏移量 映射
}

// NewConsumer 创建一个新的 Consumer 实例
func NewConsumer(cid string) *Consumer {
	return &Consumer{
		Cid:            cid,
		TopicOffsetMap: &sync.Map{},
	}
}

// SetOffset 设置特定 topic 的 offset
func (c *Consumer) SetOffset(topic string, offset int) {
	c.TopicOffsetMap.Store(topic, offset)
}

// GetOffset 获取特定 topic 的 offset
func (c *Consumer) GetOffset(topic string) (int, bool) {
	value, ok := c.TopicOffsetMap.Load(topic)
	if !ok {
		return 0, false
	}
	return value.(int), true
}

// IncrementOffset 增加特定 topic 的 offset
func (c *Consumer) IncrementOffset(topic string, increment int) {
	c.TopicOffsetMap.Store(topic, c.GetOffsetOrDefault(topic)+increment)
}

// GetOffsetOrDefault 获取特定 topic 的 offset，如果不存在则返回 0
func (c *Consumer) GetOffsetOrDefault(topic string) int {
	value, ok := c.TopicOffsetMap.Load(topic)
	if !ok {
		return 0
	}
	return value.(int)
}

// ConsumerRegister 消费者注册结构体
type ConsumerRegister struct {
	Cid   string `json:"cid"`   // 消费者ID
	Topic string `json:"topic"` // 消费者订阅的主题
}
