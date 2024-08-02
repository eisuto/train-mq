package handlers

import (
	"log"
	"net/http"
	"trainMQ/models"
	"trainMQ/queue"
	"trainMQ/utils"
)

// ConsumeHandler 消息消费处理器
func ConsumeHandler(queue *queue.MessageQueue) http.HandlerFunc {
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

		// 消费
		msg, ok := queue.Consume(topic)

		// 响应并记录日志
		clientIp := utils.GetClientIp(r)
		if ok {
			models.WriteSuccessResponse(w, "Consumed message successfully", msg)
			log.Printf("Client IP: %s - Consumed message ID: %s, Topic: %s, Content: %s\n", clientIp, msg.ID, msg.Topic, msg.Content)
		} else {
			models.WriteSuccessResponse(w, "No messages available for topic: "+topic, []int{})
			log.Printf("Client IP: %s - No messages available for topic: %s\n", clientIp, topic)
			return
		}

	}
}
