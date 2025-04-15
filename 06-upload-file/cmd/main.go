package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"

	v1 "github.com/clin211/gin-learn/06-upload-file/api/v1"
	"github.com/clin211/gin-learn/06-upload-file/internal/config"
)

func main() {
	// 初始化配置
	configPath := filepath.Join("configs", "config.yaml")
	if err := config.Init(configPath); err != nil {
		log.Fatalf("初始化配置失败: %v", err)
		return
	}

	// 现在可以通过 config.xx 来访问配置
	fmt.Println("config.Local.UploadDir: ", config.Local.UploadDir)

	// Gin 初始化
	gin.SetMode(config.Server.Mode)
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	v1.RegisterRoutes(router)

	port := fmt.Sprintf(":%d", config.Server.Port)
	fmt.Println("server start at ", port)
	router.Run(port)
}
