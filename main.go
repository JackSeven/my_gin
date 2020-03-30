package main

import (
	"fmt"
	"my_gin/config"
	"my_gin/routers"
)

func main() {
	initRouter()
}

func initRouter() {

	res := config.LoadConfig()
	if res != true {
		fmt.Println("init load config failed")
	}

	fmt.Println("load config ", res)

	// 初始化路由
	router := routers.InitRouter()
	router.Static("/static", "./static")

	// 运行
	router.Run(":6001")
}


