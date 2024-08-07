package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"train-mq/core"
	"train-mq/handlers"
)

const ConfigFileName = "train-mq-config.yaml"

func main() {
	// 读取配置文件
	config, err := LoadConfig(ConfigFileName)
	if err != nil {
		log.Printf("加载配置文件失败: %v ", err)
		return
	}

	// Banner
	if config.Banner {
		banner()
	}

	// 创建消息队列
	messageQueue := core.NewMainContext()

	log.Printf("TrainMQ launch successful! Now listening on port: %v", config.Port)
	// 开启 HTTP 服务
	http.HandleFunc("/publish", handlers.PublishHandler(messageQueue))
	http.HandleFunc("/consume", handlers.ConsumeHandler(messageQueue))
	http.HandleFunc("/subscribe", handlers.SubscribeHandler(messageQueue))
	err = http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}

}

func banner() {
	fmt.Print(`
/__  ___/                              /|    //| |     //    ) ) 
   / /   __      ___     ( )   __     //|   // | |    //    / /  
  / /  //  ) ) //   ) ) / / //   ) ) // |  //  | |   //    / /   
 / /  //      //   / / / / //   / / //  | //   | |  //  \ \ /    
/ /  //      ((___( ( / / //   / / //   |//    | | ((____\ \
`)
	fmt.Println()

}
