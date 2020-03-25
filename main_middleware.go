package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(middleWare())
	{
		router.GET("/index", func(c *gin.Context) {
			id := c.Query("id")
			c.JSON(http.StatusOK, gin.H{"ID":id})
		})
	}

	router.Run(":6001")
}

// 自定义中间件
func middleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Println("before middleware")

		c.Set("request", "client_request")
		c.Next()

		status := c.Writer.Status()
		fmt.Println("after middleware", status)

		t2 := time.Since(t)

		fmt.Println("time2:", t2)
	}
}