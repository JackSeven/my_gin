package routers

import (
	"github.com/gin-gonic/gin"
	"my_gin/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// 首页
	router.GET("/login", controllers.Index)

	// 添加搜索记录
	router.GET("/keyword", controllers.UsAdd)

	return router
}
