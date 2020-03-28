package main

import (
	"my_gin/routers"


)

func main() {
	initRouter()
}

func initRouter() {
	router := routers.InitRouter()
	router.Static("/static", "./static")
	router.Run(":6001")
}


