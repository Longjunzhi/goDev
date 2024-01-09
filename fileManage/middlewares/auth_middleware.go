package middlewares

import (
	"fileManage/constants"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			fmt.Println("c.GetHeader(constants.HEADER_KEY_BEAR_TOKEN)", c.GetHeader(constants.HEADER_KEY_BEAR_TOKEN))
			fmt.Println("c.Request.Header.Get(constants.HEADER_KEY_BEAR_TOKEN)", c.Request.Header.Get(constants.HEADER_KEY_BEAR_TOKEN))
			if c.GetHeader(constants.HEADER_KEY_BEAR_TOKEN) != constants.TEMP_TOKEN {
				c.JSON(http.StatusUnauthorized, "未授权")
				return
			}
		}
		c.Next()
	}
}
