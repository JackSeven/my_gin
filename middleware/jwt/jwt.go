package jwt

import (
	"github.com/gin-gonic/gin"
	"my_gin/libraries/myjwt"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 1
		token := c.Query("token")
		if token == "" {
			code = 0
		} else {
			claims, err := myjwt.ParseToken(token)
			if err != nil {
				code = -1
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = -2
			}
		}

		if code != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"state" : code,
				"data" : data,
				"msg" : code,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
