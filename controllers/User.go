package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clib "my_gin/libraries"
	"my_gin/models"
	"strconv"
	"time"
)

// 添加用户搜索记录
func UsAdd(c *gin.Context)  {

	keyword := c.PostForm("us")
	if keyword == "" {
		clib.ReturnError(c, "", "")
	}

	// 格式化为时间戳
	create_time,_ := strconv.Atoi(fmt.Sprintf("%v",time.Now().Unix()))
	engine_id, _ := strconv.Atoi(c.PostForm("id"))
	us := models.UserSearch{
		UserId:     1,
		Keyword:    keyword,
		EnginId:    engine_id,
		EnginName:  c.PostForm("en"),
		CreateTime: create_time,
	}

	// 插入记录
	id, err := models.UsAdd(us)
	if err != nil {
		clib.ReturnError(c, id, "faild")
	}

	clib.ReturnSuccess(c, id, "success")
	return

}

// 获取搜索历史记录
func UsList(c *gin.Context)  {

	//index := c.Query("index")
	//pageSize := c.Query("pageSize")
	//
	//



}







