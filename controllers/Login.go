package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{"title":"欢迎注册"})
}
