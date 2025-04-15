package upload

import "github.com/gin-gonic/gin"

func (u *FileUploader) Multiple(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "multiple ok",
	})
}
