package main

import (
	. "TrainMQ/entity"
	"TrainMQ/processor"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 被写入
func write(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	// 解析参数
	_ = decoder.Decode(&params)
	info := NewInfo(params["body"], int(time.Now().Unix()), 0, params["theme"])
	processor.Push(info, wayMap)
	log.Printf("【写入】主题：%s\t 信息：%s\t 写入时间：%d\n", params["theme"], params["body"], time.Now().Unix())

	reJson, _ := json.Marshal(info)
	_, _ = fmt.Fprintf(writer, string(reJson))
}

// 被读出
func read(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]string
	_ = decoder.Decode(&params)
	theme := params["theme"]
	po := processor.Pop(theme, wayMap)
	po.ReceiveTime = int(time.Now().Unix())
	log.Printf("【读出】主题：%s\t 信息：%s\t 读出时间：%d\n", params["theme"], po.Body, po.ReceiveTime)
	
	reJson, _ := json.Marshal(po)
	_, _ = fmt.Fprintf(writer, string(reJson))

}
