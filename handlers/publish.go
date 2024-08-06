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

// PublishHandler 发布请求处理
func PublishHandler(queue *core.MainMessageQueue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		var msg models.Message
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body", nil)
			return
		}
		// 生成唯一ID
		msg.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
		// 发布消息
		queue.Publish(msg)
		// 响应并记录日志
		clientIP := utils.GetClientIp(r)
		models.WriteSuccessResponse(w, "Published message successfully", msg)
		log.Printf("Client IP: %s - Published message ID: %s, Topic: %s, Content: %s\n", clientIP, msg.ID, msg.Topic, msg.Content)

	}
}
