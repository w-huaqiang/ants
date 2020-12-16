package main

import (
	"bjzdgt.com/ants/controller"
	"bjzdgt.com/ants/middleware"
	"github.com/gin-gonic/gin"
)

// CollectRoute for route
func CollectRoute(r *gin.Engine) *gin.Engine {

	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)

	return r
}
