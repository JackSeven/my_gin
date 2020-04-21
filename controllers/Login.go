package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clib "my_gin/libraries"
	"my_gin/libraries/myjwt"
	"net/http"
)

// 首页
func Index(c *gin.Context)  {

	// 开始时间
	start := clib.ExecTime("start", 0)

	// 获取mac地址
	macAddr := clib.GetMacAddr()
	fmt.Println(macAddr)

	c.HTML(http.StatusOK, "index.html", gin.H{"title":"夸克导航"})

	// 结束时间
	clib.ExecTime("end", start)

}

// 获取授权token
func Login(c *gin.Context)  {


	param := map[string]string {
		"username": "admin",
		"password":"123456",
	}

	token := myjwt.GenerateToken(param)

	fmt.Println("get token...................")
	fmt.Println(token)
	fmt.Println("parse token...................")

	info,err := myjwt.ParseToken(token)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
	fmt.Println(info.Username)
	fmt.Println(info.Password)
	fmt.Println(info.ExpiresAt)

	clib.ReturnSuccess(c, token,"success")

}


