package controller

import (
	"errors"
	"fileManage/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	ctx := c.Request.Context()
	req := &services.LoginByPasswordRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, code, err := services.LoginByPassword(ctx, req)
	if code != http.StatusOK {
		if err != nil {
			c.JSON(code, err.Error())
			return
		}
		c.JSON(code, errors.New("no err message"))
		return
	}
	if err != nil {
		c.JSON(code, err.Error())
		return
	}
	c.JSON(http.StatusOK, "resp")
}
