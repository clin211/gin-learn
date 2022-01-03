package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type News struct {
	Title   string
	Content string
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("pages/**/*")

	// layout
	r.GET("/layout", func(c *gin.Context) {
		c.HTML(http.StatusOK, "layout/index.html", gin.H{
			"title": "layout home",
		})
	})

	r.GET("/layout/news", func(c *gin.Context) {
		news := &News{
			Title:   "layout new title",
			Content: "layout new content",
		}

		c.HTML(http.StatusOK, "layout/news.html", gin.H{
			"title": "news",
			"news":  news,
		})
	})

	// admin
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "admin home",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		news := &News{
			Title:   "admin new title",
			Content: "admin new content",
		}

		c.HTML(http.StatusOK, "layout/news.html", gin.H{
			"title": "news",
			"news":  news,
		})
	})

	// default
	r.GET("/default", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "default home",
			"data": []string{"javascript", "node", "vue", "react", "next", "nuxt", "react-native", "flutter"},
		})
	})

	r.GET("/default/news", func(c *gin.Context) {
		news := &News{
			Title:   "default new title",
			Content: "default new content",
		}

		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "news",
			"news":  news,
		})
	})

	r.Run(":9090")
}
