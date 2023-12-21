package routes

import (
	"Img/config"
	"Img/controller"
	"Img/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	Routes *gin.Engine
)

func init() {
	Routes = gin.Default()
	r := Routes.Use(middlewares.AuthMiddleWare(), middlewares.Cors())
	r.Static("/public", config.AppConf.StorageConf.Path)
	r.POST("/api/login", controller.Login)
	r.POST("/api/media/upload", controller.Upload)
	r.POST("/api/media/upload/multiple", controller.UploadMultiple)
	r.POST("/api/media/get", controller.MediaGet)
}
