package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 获取默认服务器
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world")
	})

	// get 参数
	//router.GET("/get/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.String(http.StatusOK, name)
	//  // 返回结果 输入"name"的值
	//})

	// router1, get 参数
	router.GET("/get/:name/*action", func(c *gin.Context) {
		action := c.Param("action")
		name := c.Param("name")
		c.String(http.StatusOK, name)
		c.String(http.StatusOK, action)

		println(name)
		println(action)

		// http://127.0.0.1:6001/get/a/b
		// 返回结果 "a/b"
	})

	// router2, query 接受参数
	router.GET("/get", func(c *gin.Context) {
		// 设置接受query的key 和 默认值
		id := c.DefaultQuery("id", "")
		fmt.Println(id)

		println(id)

		c.String(http.StatusOK, fmt.Sprintf("id is: %s ",id))
		// http://127.0.0.1:6001/get?id=2323
		// 返回结果 ： id is: 2323
	})

	// router3, postform 表单提交 view/submit.html
	router.POST("/submit", func(c *gin.Context) {
		username := c.PostForm("username")
		userpwd  := c.PostForm("userpwd")
		remeber  := c.DefaultPostForm("remeber", "1")
		hobby    := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf(
			"接受到的参数为：" +
			"username: %s" +
			"userpwd: %s" +
			"remember:%s" +
			"hobby:%v" +
		username, userpwd, remeber, hobby))
	})



	// 启动方式1
	router.Run(":6001")

	// 启动方式2
	//http.ListenAndServe(":6001", router)

	// 启动方式3
	//s := &http.Server{
	//	Addr : "6001",
	//	Handler : router,
	//	ReadTimeout : 10 * time.Second,
	//	WriteTimeout : 10 * time.Second,
	//	MaxHeaderBytes : 1 << 20,
	//}
	//
	//s.ListenAndServe()

}

