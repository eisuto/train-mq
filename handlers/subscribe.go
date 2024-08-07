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
func SubscribeHandler(context *core.MainContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取参数
		var consumerRegister models.ConsumerRegister
		if err := json.NewDecoder(r.Body).Decode(&consumerRegister); err != nil {
			models.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", nil)
			return
		}
		// 主题不存在不能订阅
		_, ok := context.GetQueue(consumerRegister.Topic)
		if !ok {
			models.WriteErrorResponse(w, http.StatusNotFound, "Topic does not exist", nil)
			log.Printf("Topic: %s does not exist, created and registered consumerRegister: %v", consumerRegister.Topic, consumerRegister)
			return
		}
		// 生成唯一ID
		consumerRegister.Cid = strings.ReplaceAll(uuid.NewString(), "-", "")
		// 注册消费者
		context.RegisterConsumer(consumerRegister.Topic, consumerRegister.Cid)
		// 响应并记录日志
		clientIP := utils.GetClientIp(r)
		models.WriteSuccessResponse(w, "Registered consumerRegister successfully", consumerRegister)
		log.Printf("Client IP: %s - Registered consumerRegister ID: %s, Topic: %s\n", clientIP, consumerRegister.Cid, consumerRegister.Topic)

	}
}
