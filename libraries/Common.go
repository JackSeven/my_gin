package libraries

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


// 返回成功json
func ReturnSuccess(c *gin.Context, data interface{}, msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"state" : true,
		"data" : data,
		"msg" : msg,
	})
	return
}

// 返回失败json
func ReturnError(c *gin.Context, data interface{}, msg string)  {
	c.JSON(http.StatusOK, gin.H{
		"state" : false,
		"data" : data,
		"msg" : msg,
	})
	return
}