package controllers

import (
	"gateway/responses"
	// "gateway/configs"
	// "fmt"
	"log"
	"net/http"
	// "time"

	// pb "example.com/blogArch/proto"
	// "google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
// var validate = validator.New()

func GetMainPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In main controller...")

		resp := responses.MainResponse{
			Output: "In Main Page!",
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In profile controller...")

		resp := responses.ProfileResponse{
			UserID:  "ID: John Doe",
			Entries: []string{"FirstEntry", "SecondEntry"},
		}
		c.JSON(http.StatusOK, resp)
	}
}

func Entry() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In entry controller...")

		resp := responses.EntryResponse{
			Status: "Entry Posted!",
		}
		c.JSON(http.StatusOK, resp)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In login controller...")

		resp := responses.LoginResponse{
			User: "User",
			Password: "Password",
		}
		c.JSON(http.StatusOK, resp)
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In register controller...")

		resp := responses.RegisterResponse{
			User: "RegisterUser",
			Password: "RegisterPassword".
		}
		c.JSON(http.StatusOK, resp)
	}
}
