package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	//router.LoadHTMLGlob("view/*")

	router.LoadHTMLFiles("view/index.html")

	router.GET("/index", func(c *gin.Context) {
		id := c.Query("id")
		if id == "12" {
			c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "文章标题",
			"content": fmt.Sprintf("文章内容%s", id),
		})
	})

	router.Run(":6001")

}
