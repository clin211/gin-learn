package v1

import (
	"github.com/clin211/gin-learn/06-upload-file/api/v1/upload"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	uploader := &upload.FileUploader{}

	api := router.Group("/api/v1")
	{
		// 文件上传
		uploadRouter := api.Group("/upload")
		{
			uploadRouter.POST("/file", uploader.File)               // 上传单文件
			uploadRouter.POST("/multiple", uploader.Multiple)       // 上传多个文件
			uploadRouter.POST("/folder", uploader.Folder)           // 上传文件夹
			uploadRouter.POST("/chunk/init", uploader.ChunkInit)    // 初始化分片上传，返回 uploadID
			uploadRouter.POST("/chunk", uploader.Chunk)             // 上传分片（服务端自动触发合并）
			uploadRouter.GET("/chunk/status", uploader.ChunkStatus) // 查询分片上传状态（可选，用于断点续传）
			uploadRouter.GET("/url", uploader.GetFileURL)           // 获取文件访问地址（支持 MinIO 对象存储）
			uploadRouter.GET("/meta/:file_name", uploader.Meta)     // 获取文件元信息（如大小、类型等）
			uploadRouter.DELETE("/:file_name", uploader.Delete)     // 删除文件（通过 file_name 或 object_key）
		}
	}
}
