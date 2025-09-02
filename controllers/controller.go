package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// 可以导入 models
	// "your-gin-app/models"
)

// HomePage 示例控制器
func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Gin App!")
}

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {
	// 实际应用中，这里会调用模型层获取数据
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

// CreateUser 创建新用户
func CreateUser(c *gin.Context) {
	// 实际应用中，这里会绑定请求体并调用模型层处理
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
