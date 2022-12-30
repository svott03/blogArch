package routes

import (
	"gateway/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controllers.GetMainPage())
	router.GET("/profile/:userId", controllers.GetProfile())
	router.POST("/entry", controllers.Entry())
	router.POST("/login", controllers.Login())
	router.POST("/register", controllers.Register())
}