package controllers

import (
	"github.com/gin-gonic/gin"
	clib "my_gin/libraries"
	"my_gin/models"
)

// 引擎列表
func Lists(c *gin.Context)  {
	clib.ReturnSuccess(c, models.Lists(), "success")
}
