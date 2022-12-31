package controllers

import (
	"example.com/blogArch/gateway/models"
	"example.com/blogArch/gateway/responses"
	"example.com/blogArch/gateway/configs"
	"example.com/blogArch/gateway/utils"
	"context"
	"log"
	"net/http"
	"time"

	pb "example.com/blogArch/proto"
	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)


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
		// TODO add JWT for passing authentication information
		resp := responses.ProfileResponse{
			UserID:  "ID: John Doe",
			Entries: []string{"FirstEntry", "SecondEntry"},
		}
		c.JSON(http.StatusOK, resp)
	}
}

// Add gRPC data models
const (
	ADDRESS = "localhost:50051"
)

type FilterTask struct {
	Input string
}

func Entry() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In entry controller...")
		// TODO JWT authentication information
		var body models.EntryModel
		c.BindJSON(&body)
		// Send information to gRPC
		conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

		if err != nil {
			log.Fatalf("did not connect : %v", err)
		}

		defer conn.Close()

		// Establish connection with microservice
		cont := pb.NewTextFilterServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 80*time.Second)

		defer cancel()
		res, err := cont.CreateFilterOutput(ctx, &pb.FilterInput{Input: body.Entry})
		
		log.Println("Ouput is %s", res.GetOutput())

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		// TODO entry positive/negative cases, insert into db then send appropriate response msg
		resp := responses.StatusResponse{
			Status: res.GetOutput(),
		}
		c.JSON(http.StatusOK, resp)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In login controller...")
		var body models.LoginModel
		c.BindJSON(&body)

		status := utils.TryLogin(body.Username, body.Password)

		resp := responses.StatusResponse{
			Status: status,
		}
		// TODO return JWT?
		c.JSON(http.StatusOK, resp)
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In register controller...")
		var body models.RegisterModel
		c.BindJSON(&body)
		status := utils.TryRegister(body.Username, body.Password)
		resp := responses.StatusResponse{
			Status: status,
		}
		// TODO return JWT?
		c.JSON(http.StatusOK, resp)
	}
}
