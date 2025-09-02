package routes

import (
	"gin_app/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes 定义所有路由
func SetupRoutes(router *gin.Engine) {
	// 根路径
	router.GET("/", controllers.HomePage)

	// API 路由组
	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/users", controllers.GetUsers)
		apiGroup.POST("/users", controllers.CreateUser)
	}
}
