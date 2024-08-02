package models

import (
	"encoding/json"
	"net/http"
)

const (
	success            = 200  // 正常响应
	Error              = 500  // 业务异常
	ParameterError     = 400  // 参数错误
	InvalidRequestBody = 1001 // 请求正文无效
	NoMessage          = 2001 // 无消息

)

// Response 通用响应
type Response struct {
	Code    int         `json:"code"`           // 响应码
	Message string      `json:"message"`        // 信息
	Data    interface{} `json:"data,omitempty"` // 数据
}

// WriteSuccessResponse 写入成功响应
func WriteSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	if message == "" {
		message = "ok"
	}
	response := Response{
		Code:    success,
		Message: message,
		Data:    data,
	}
	writeResponse(w, response)
}

// WriteErrorResponse 写入失败响应
func WriteErrorResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	if message == "" {
		message = "error"
	}
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	writeResponse(w, response)
}

// writeResponse 写入响应
func writeResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
