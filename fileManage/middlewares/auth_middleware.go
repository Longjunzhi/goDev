package middlewares

import (
	"Img/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, ok := c.Get("Authorization"); ok {
			fmt.Println("Authorization", token)
		}
		fmt.Println("Header Authorization", c.GetHeader("Authorization"))
		if token, ok := c.Get(constants.HEADER_KEY_BEAR_TOKEN); ok {
			if token != constants.TEMP_TOKEN {
				c.JSON(http.StatusUnauthorized, "未授权")
				return
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
