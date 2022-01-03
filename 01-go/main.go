package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 路由
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "获取用户信息",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "创建一条用户信息",
		})
	})

	r.PUT("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "更新一条用户信息",
		})
	})

	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": "删除一条用户信息",
		})
	})

	// 监听并在 0.0.0.0:8080 上启动服务，可以传一个监听端口的字符串 r.Run("9090")
	r.Run(":9090")
}
