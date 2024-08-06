package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"train-mq/core"
	"train-mq/models"
	"train-mq/utils"
)

// SubscribeHandler 订阅主题处理
func SubscribeHandler(queue *core.MainMessageQueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取参数
		var consumer models.Consumer
		if err := json.NewDecoder(r.Body).Decode(&consumer); err != nil {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			return
		}
		// 主题不存在不能订阅
		_, ok := queue.GetQueue(consumer.Topic)
		if !ok {
			models.WriteErrorResponse(w, http.StatusNotFound, "Topic does not exist", nil)
			log.Printf("Topic: %s does not exist, created and registered consumer: %v", consumer.Topic, consumer)
			return
		}
		// 生成唯一ID
		consumer.Cid = strings.ReplaceAll(uuid.NewString(), "-", "")
		// 注册消费者
		queue.RegisterConsumer(consumer.Topic, consumer)
		// 响应并记录日志
		clientIP := utils.GetClientIp(r)
		models.WriteSuccessResponse(w, "Registered consumer successfully", consumer)
		log.Printf("Client IP: %s - Registered consumer ID: %s, Topic: %s\n", clientIP, consumer.Cid, consumer.Topic)

	}
}
