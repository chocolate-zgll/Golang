package config

import (
	"os"
)

type Config struct {
	Port string
	// 可以添加其他配置，如数据库连接字符串等
	// DB_URI string
}

// LoadConfig 从环境变量或默认值加载配置
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	return &Config{
		Port: port,
		// DB_URI: os.Getenv("DB_URI"),
	}
}
