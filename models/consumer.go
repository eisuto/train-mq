package models

// Consumer 消费者结构体
type Consumer struct {
	Cid    string   `json:"cid"`    // 消费者ID
	Topics []string `json:"topics"` // 消费者订阅的主题
	Offset int64    `json:"offset"` // 消费者的偏移量，记录了上次消费的位置
}

// ConsumerRegister 消费者注册结构体
type ConsumerRegister struct {
	Cid   string `json:"cid"`   // 消费者ID
	Topic string `json:"topic"` // 消费者订阅的主题
}
