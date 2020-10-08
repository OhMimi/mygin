package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 說Hello
// @Id 1
// @Tags Hello
// @version 1.0
// @produce text/plain
// @Success 200 string string 成功後返回的值
// @Router /hello [get]
func GetHello(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("Hello, It Home!"))
}

// @Summary Delete Hello
// @Id 1
// @Tags Hello
// @version 1.0
// @produce text/plain
// @param id path int true "id"
// @Success 200 string string 成功後返回的值
// @Router /hello/{id} [delete]
func DeleteHello(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "hello DELETE %s", id)
}
