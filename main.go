package main

import (
	"github.com/gin-gonic/gin"
	"github.com/evilGopher/controller"
	"net/http"
)

func router() http.Handler {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.GET("/health-check", controller.HealthCkeck)

	r.POST("/user/register", controller.RegisterUser)
	r.POST("/user/login", controller.LoginUser)
	r.POST("/user/logout", controller.LogoutUser)


	return r
}

func main() {
	r := &http.Server{
		Addr:           ":8080",
		Handler:        router(),
	}
	r.ListenAndServe()
}
