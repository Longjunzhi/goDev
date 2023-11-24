package controller

import (
	"Img/config"
	"Img/util"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	open, err := file.Open()
	if err != nil {
		return
	}
	hash := md5.New()
	_, _ = io.Copy(hash, open)
	md5String := hex.EncodeToString(hash.Sum(nil))
	err = c.SaveUploadedFile(file, config.AppConf.BasePath+util.GetPathTag()+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, md5String)
}

func UploadMultiple(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	files := form.File["files"]
	var md5String []string
	for _, file := range files {
		open, err := file.Open()
		if err != nil {
			return
		}
		hash := md5.New()
		_, _ = io.Copy(hash, open)
		md5String = append(md5String, hex.EncodeToString(hash.Sum(nil)))
		if err := c.SaveUploadedFile(file, config.AppConf.BasePath+util.GetPathTag()+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, md5String)
}
