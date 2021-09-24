package main

import (
	. "TrainMQ/entity"
	"fmt"
	"log"
	"net/http"
	"time"
)

var wayMap map[string]Queue
var queue Queue

//var wayMap sync.Map

func main() {
	wayMap = make(map[string]Queue)
	fmt.Println(`
      .---- -. -. .  .   .        .                                           
     ( .',----- - - ' '      '                                         __     
      \_/      ;--:-          __--------------------___  ____=========_||___  
     __U__n_^_''__[.  ooo___  | |_!_||_!_||_!_||_!_| |   |..|_i_|..|_i_|..|   
   c(_ ..(_ ..(_ ..( /,,,,,,] | |___||___||___||___| |   |                |   
   ,_\___________'_|,L______],|______________________|_i,!________________!_i 
  /;_(@)(@)==(@)(@)   (o)(o)      (o)^(o)--(o)^(o)          (o)(o)-(o)(o)     
""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""~"""

      /__  ___/                              /|    //| |     //    ) ) 
         / /   __      ___     ( )   __     //|   // | |    //    / /  
        / /  //  ) ) //   ) ) / / //   ) ) // |  //  | |   //    / /   
       / /  //      //   / / / / //   / / //  | //   | |  //  \ \ /    
      / /  //      ((___( ( / / //   / / //   |//    | | ((____\ \

启动成功! 当前监听端口 :5757
开始享受您如同高速铁路般的信息传递之旅吧！`)

	s := new(http.Server)
	s.ReadTimeout = 5 * time.Second
	s.WriteTimeout = 5 * time.Second
	http.HandleFunc("/write", write)
	http.HandleFunc("/read", read)
	err := http.ListenAndServe(":5757", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
