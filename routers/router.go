package routers

import (
	"github.com/gin-gonic/gin"
	"my_gin/controllers"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/login", controllers.Index)
	return router
}
