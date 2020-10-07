package handler

import (
	"mygin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	c.JSON(http.StatusOK, gin.H{
		"userId": uid,
	})
}

func UserLogin(c *gin.Context) {
	var user model.UserLogin
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
