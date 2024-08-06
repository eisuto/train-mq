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
		var subscriber models.Subscriber
		if err := json.NewDecoder(r.Body).Decode(&subscriber); err != nil {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			return
		}
		// 主题不存在不能订阅
		_, ok := queue.GetQueue(subscriber.Topic)
		if !ok {
			models.WriteErrorResponse(w, http.StatusNotFound, "Topic does not exist", nil)
			log.Printf("Topic: %s does not exist, created and registered subscriber: %v", subscriber.Topic, subscriber)
			return
		}
		// 生成唯一ID
		subscriber.SubId = strings.ReplaceAll(uuid.NewString(), "-", "")
		// 注册订阅者
		queue.RegisterSubscriber(subscriber.Topic, subscriber)
		// 响应并记录日志
		clientIP := utils.GetClientIp(r)
		models.WriteSuccessResponse(w, "Registered subscriber successfully", subscriber)
		log.Printf("Client IP: %s - Registered subscriber ID: %s, Topic: %s\n", clientIP, subscriber.SubId, subscriber.Topic)

	}
}
