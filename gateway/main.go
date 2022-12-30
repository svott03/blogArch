package main

import (
	"gateway/configs"
	"gateway/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// configs.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}