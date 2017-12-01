package main

import (
	"github.com/gin-gonic/gin"
	"github.com/evilGopher/controller"
)

func main() {
	r := gin.Default()
	r.GET("/health-check", controller.HealthCkeck)
	r.Run() // listen and serve on 0.0.0.0:8080
}
