package routes

import (
	"example.com/blogArch/gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.GetMainPage())
	// TODO jwt username tokens
	router.GET("/profile", controllers.GetProfile())
	router.POST("/entry", controllers.Entry())
	router.POST("/login", controllers.Login())
	router.POST("/register", controllers.Register())
}