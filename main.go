package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type News struct {
	Title   string
	Content string
}

// 时间戳转换日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	fmt.Print(t)
	return t.Format("2006-01-02 15:04:05")
}

func main() {

	r := gin.Default()
	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
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
		fmt.Print(UnixToTime(int(time.Now().Unix())))
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "default home",
			"data":  []string{"javascript", "node", "vue", "react", "next", "nuxt", "react-native", "flutter"},
			"time":  1641204239, //int64(time.Now().Unix()),
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
