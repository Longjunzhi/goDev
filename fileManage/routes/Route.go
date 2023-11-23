package routes

import (
	"Img/controller"
	"github.com/gin-gonic/gin"
)

var (
	Routes *gin.Engine
)

func init() {
	Routes = gin.Default()
	Routes.POST("/api/login", controller.Login)
	Routes.POST("/api/upload", controller.Upload)
	Routes.POST("/api/upload/multiple", controller.UploadMultiple)
}
