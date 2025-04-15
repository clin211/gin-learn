package upload

import "github.com/gin-gonic/gin"

func (u *FileUploader) Folder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "folder ok",
	})
}
