package main

import (
	"example.com/blogArch/gateway/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run("localhost:8080")
}