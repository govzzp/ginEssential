package routes

import (
	"ginEsseential/controller"
	"ginEsseential/middleware"
	"github.com/gin-gonic/gin"
)

func ConnectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register",controller.Register )
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)
	return r
}