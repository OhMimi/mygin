package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("Hello, It Home!"))
}

func DeleteHello(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "hello DELETE %s", id)
}
