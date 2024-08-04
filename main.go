package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"train-mq/handlers"
	"train-mq/queue"
)

const ConfigFileName = "train-mq-config.yaml"

func main() {
	// 读取配置文件
	config, err := LoadConfig(ConfigFileName)
	if err != nil {
		log.Printf("加载配置文件失败: %v ", err)
		return
	}
	port := config.Port
	banner(port)

	// 创建消息队列
	messageQueue := queue.NewMessageQueue()

	// 开启 HTTP 服务
	http.HandleFunc("/publish", handlers.PublishHandler(messageQueue))
	http.HandleFunc("/consume", handlers.ConsumeHandler(messageQueue))
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func banner(port int) {
	fmt.Printf(`
/__  ___/                              /|    //| |     //    ) ) 
   / /   __      ___     ( )   __     //|   // | |    //    / /  
  / /  //  ) ) //   ) ) / / //   ) ) // |  //  | |   //    / /   
 / /  //      //   / / / / //   / / //  | //   | |  //  \ \ /    
/ /  //      ((___( ( / / //   / / //   |//    | | ((____\ \
Launch successful! Now listening on port: %v`, port)
	fmt.Println()

}
