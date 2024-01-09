package controller

import (
	"errors"
	"fileManage/constants"
	"fileManage/services"
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
	mediaGetRes, code, err := services.MediaGet(ctx, req)
	if code != http.StatusOK {
		if err != nil {
			c.JSON(code, err.Error())
			return
		}
		c.JSON(code, errors.New("no err message"))
		return
	}
	for i, m := range mediaGetRes.Media {
		mediaGetRes.Media[i].Path = constants.APP_STORAGE_URL + m.Path
	}
	resp.Data = mediaGetRes
	c.JSON(http.StatusOK, resp)
}
