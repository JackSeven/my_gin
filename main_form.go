package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//type LoginInfo struct {
//	Username string	`json:"username"`
//	Userpwd string `json:"userpwd"`
//}

//数据绑定 接收json 类型数据

//func main() {
//
//	router := gin.Default()
//
//	router.POST("/submit", func(c *gin.Context) {
//		var json LoginInfo
//		// 接受json , 用postman 的 json 类型调试
//		if err := c.ShouldBindJSON(&json); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if json.Username != "admin" || json.Userpwd != "123456" {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "", "msg":"success"})
//	})
//
//	router.Run(":6001")
//}
//


// 数据绑定 接收 form 表单提交数据

type LoginInfo struct {
	Username string	`form:"username"`
	Userpwd string `form:"userpwd"`
}

func main() {
	router := gin.Default()

	router.POST("/submit", func(c *gin.Context) {

		var form LoginInfo
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{"error":err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"username": form.Username, "userpwd": form.Userpwd})
		
	})

	router.Run(":6001")

}