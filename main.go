package main

import (
	"fmt"
	"log"
	"net/http"
	"trainMQ/handlers"
	"trainMQ/queue"
)

//var wayMap sync.Map

func main() {
	fmt.Println(`
/__  ___/                              /|    //| |     //    ) ) 
   / /   __      ___     ( )   __     //|   // | |    //    / /  
  / /  //  ) ) //   ) ) / / //   ) ) // |  //  | |   //    / /   
 / /  //      //   / / / / //   / / //  | //   | |  //  \ \ /    
/ /  //      ((___( ( / / //   / / //   |//    | | ((____\ \

Launch successful! Now listening on port: 5757`)

	messageQueue := queue.NewMessageQueue()

	http.HandleFunc("/publish", handlers.PublishHandler(messageQueue))
	http.HandleFunc("/consume", handlers.ConsumeHandler(messageQueue))

	err := http.ListenAndServe(":5757", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
