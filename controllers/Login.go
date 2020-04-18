package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clib "my_gin/libraries"
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



