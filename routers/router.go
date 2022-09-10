package routers

import (
	"net/http"

	"github.com/Forest-211/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "pong",
		})
	})
	return r
}
