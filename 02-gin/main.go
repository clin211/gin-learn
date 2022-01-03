package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	title       string
	description string
	content     string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pages/*")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "%v page", "home")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ok",
			"data":    "json1",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "ok",
			"data":    "json2",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		article := &Article{
			title:       "gin",
			content:     "content",
			description: "description",
		}
		fmt.Println(article)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "ok",
			"data":    article,
		})
	})

	//返回xml数据
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": true,
			"message": "返回xml",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title":   "this is news",
			"content": "gin html 模板渲染",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title":   "this is news",
			"content": "iPad mini Mega power. Mini sized.",
		})
	})

	r.Run(":9090")
}
