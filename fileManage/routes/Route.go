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
	Routes.Use(middlewares.AuthMiddleWare())
	Routes.Use(middlewares.Cors())
	Routes.Static("/public", config.AppConf.StorageConf.Path)
	Routes.POST("/api/login", controller.Login)
	Routes.POST("/api/media/upload", controller.Upload)
	Routes.POST("/api/media/upload/multiple", controller.UploadMultiple)
	Routes.POST("/api/media/get", controller.MediaGet)
}
