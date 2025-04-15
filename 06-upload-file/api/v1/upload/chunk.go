package upload

import "github.com/gin-gonic/gin"

func (u *FileUploader) Chunk(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "chunk ok",
	})
}
