package main

import (
	"example.com/blogArch/gateway/routes"
	"example.com/blogArch/gateway/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.SetupDB()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}