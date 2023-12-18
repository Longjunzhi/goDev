package controller

import (
	"Img/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MediaGet(c *gin.Context) {
	resp := services.NewSuccessResponse()
	ctx := c.Request.Context()
	req := &services.MediaGetRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	media, code, err := services.MediaGet(ctx, req)
	if code != http.StatusOK {
		if err != nil {
			c.JSON(code, err.Error())
			return
		}
		c.JSON(code, errors.New("no err message"))
		return
	}
	resp.Data = media
	c.JSON(http.StatusOK, resp)
}
