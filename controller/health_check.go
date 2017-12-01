package controller

import "github.com/gin-gonic/gin"

func HealthCkeck(c *gin.Context) {
	c.JSON(200, gin.H{
		"app": "evilGopher",
	})
}