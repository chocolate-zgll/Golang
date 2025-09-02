package main

import (
	"gin_app/config"
	"gin_app/routes"
	"github.com/gin-gonic/gin"
	"log"

	
)

func main() {
	//fmt.Println("Hello World")
	// 1. 初始化配置
	cfg := config.LoadConfig()

	// 2. 创建 Gin 引擎
	r := gin.Default()

	// 3. 注册路由
	routes.SetupRoutes(r)

	// 4. 运行服务器
	log.Printf("Server is running on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
