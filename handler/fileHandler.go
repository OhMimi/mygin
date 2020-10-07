package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadSingleIndex(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = c.SaveUploadedFile(file, "file/"+"demo.png")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileName": file.Filename,
		"size":     file.Size,
		"mimeType": file.Header,
	})
}
