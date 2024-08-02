package models

import (
	"encoding/json"
	"net/http"
)

// Response 通用响应
type Response struct {
	Message string      `json:"message"`        // 信息
	Data    interface{} `json:"data,omitempty"` // 数据
}

// WriteSuccessResponse 写入成功响应
func WriteSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	if message == "" {
		message = "ok"
	}
	response := Response{
		Message: message,
		Data:    data,
	}
	writeResponse(w, http.StatusOK, response)
}

// WriteErrorResponse 写入失败响应
func WriteErrorResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	if message == "" {
		message = "error"
	}
	response := Response{
		Message: message,
		Data:    data,
	}
	writeResponse(w, status, response)
}

// writeResponse 写入响应
func writeResponse(w http.ResponseWriter, status int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
