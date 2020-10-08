package handler

import (
	"mygin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Info
// @Tags User
// @version 1.0
// @produce text/plain
// @param Authorization header string true "Authorization"
// @param uid path int true "uid"
// @Success 200 string string 成功後返回的值
// @Router /user/{uid} [get]
func GetUser(c *gin.Context) {
	uid := c.Param("uid")
	c.JSON(http.StatusOK, gin.H{
		"userId": uid,
	})
}

// @Summary User Login
// @Tags User
// @version 1.0
// @produce application/json
// @param email formData string true "email"
// @param password formData string true "password"
// @param password-again formData string true "password-again"
// @Success 200 string string 成功後返回的值
// @Router /user/login [post]
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
