package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/evilGopher/domain"
	"github.com/evilGopher/service/user"
)

var userService user.Service

func RegisterUser(c *gin.Context) {

	var user domain.User

	if err := c.ShouldBindJSON(&user); err == nil {
		err := userService.RegisterUser(&user)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.Writer.WriteHeader(http.StatusCreated)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	return
}

