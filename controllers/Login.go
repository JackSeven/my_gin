package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_gin/models"
	"net/http"
)

func Index(c *gin.Context)  {

	if err:=models.InitDatabase(); err!=true {
		fmt.Println("connetc db faild", err)
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{"title":"欢迎注册"})
}

