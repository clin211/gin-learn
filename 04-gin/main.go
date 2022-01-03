package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/*")

	// GET 请求传值
	r.GET("/", func(c *gin.Context) {
		// Query() 获取请求时携带的参数 DefaultQuery()如果没有获取到携带的参数则第二个参数为默认值
		username := c.Query("username")
		email := c.Query("email")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"data": map[string]interface{}{
				"username": username,
				"email":    email,
				"page":     page,
			},
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{})
	})

	r.POST("/add-user", func(c *gin.Context) {
		// 获取表单传递的数据
		username := c.PostForm("username")
		email := c.PostForm("email")
		hobby := c.DefaultPostForm("hobby", "react")
		fmt.Println(username, email)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "ok",
			"dsata": map[string]interface{}{
				"username": username,
				"email":    email,
				"hobby":    hobby,
			},
		})
	})

	r.Run(":9090")
}
