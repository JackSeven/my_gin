package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页
func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{"title":"夸克导航"})
}



