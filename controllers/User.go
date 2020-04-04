package controllers

import (
	"github.com/gin-gonic/gin"
	clib "my_gin/libraries"
	"my_gin/models"
)

// 添加用户搜索记录
func UsAdd(c *gin.Context)  {

	keyword := c.Query("us")
	if keyword == "" {
		clib.ReturnError(c, "", "")
	}

	us := models.UserSearch{
		UserId:     1,
		Keyword:    keyword,
		EnginId:    1,
		EnginName:  "baidu",
		CreateTime: 12123, // 时间戳转int //todo
	}

	// 插入记录
	id, err := models.UsAdd(us)
	if err != nil {
		clib.ReturnError(c, id, "faild")
	}

	clib.ReturnSuccess(c, id, "success")
	return

}


