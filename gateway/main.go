package main

import (
	"example.com/blogArch/gateway/routes"
	"example.com/blogArch/gateway/configs"
	"example.com/blogArch/gateway/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.SetupDB()
	routes.UserRoute(router)

	protected := router.Group("/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	router.Run("localhost:8080")
}