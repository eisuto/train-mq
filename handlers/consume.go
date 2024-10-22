package handlers

import (
	"log"
	"net/http"
	"train-mq/core"
	"train-mq/models"
	"train-mq/utils"
)

// ConsumeHandler 消息消费处理器
func ConsumeHandler(content *core.MainContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// 获取主题参数
		topic := r.URL.Query().Get("topic")
		if topic == "" {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Topic parameter is missing", nil)
			return
		}
		cid := r.URL.Query().Get("cid")
		if cid == "" {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Cid parameter is missing", nil)
			return
		}
		// 只有订阅了当前主题，才能消费
		hasConsumer := content.HasConsumer(topic, cid)
		if !hasConsumer {
			models.WriteErrorResponse(w, http.StatusForbidden, "Consumer is not registered for this topic", nil)
			log.Printf("Consumer: %s is not registered for topic: %s", cid, topic)
			return
		}
		consumer := content.GetConsumerByCid(cid)
		offset, _ := consumer.TopicOffsetMap.Load(topic)
		
		// 消费
		msg, ok := content.Consume(topic, offset.(int))

		// 消费后 offset+1
		consumer.IncrementOffset(topic, 1)

		// 响应并记录日志
		clientIp := utils.GetClientIp(r)
		if ok {
			models.WriteSuccessResponse(w, "Consumed message successfully", msg)
			log.Printf("Client IP: %s - Consumed  message ID: %s, Topic: %s, offset: %v\n", clientIp, msg.ID, msg.Topic, offset)
		} else {
			models.WriteSuccessResponse(w, "No messages available for topic: "+topic, []int{})
			log.Printf("Client IP: %s - No messages available for topic: %s\n", clientIp, topic)
			return
		}

	}
}
