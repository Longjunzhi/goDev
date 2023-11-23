package controller

import (
	"Img/config"
	"Img/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = c.SaveUploadedFile(file, config.AppConf.BasePath+util.GetPathTag()+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "success")
}
func UploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	files := form.File["files"]
	for _, file := range files {
		if err := c.SaveUploadedFile(file, config.AppConf.BasePath+util.GetPathTag()+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "success")
}
