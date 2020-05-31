package main

import (
	m "ki-sdk/model"
	r "ki-sdk/router"
	"log"
)

func main() {

	// App
	m.InitSDK()

	// 初始化路由
	egg := r.InitRouter()

	// 启动  server
	err := egg.Run(":80")
	if err == nil {
		log.Println("egg is starting")
	} else {
		log.Println("egg is err:", err)
	}
}
