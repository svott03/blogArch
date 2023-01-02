package controllers

import (
	"context"
	"example.com/blogArch/gateway/models"
	"example.com/blogArch/gateway/responses"
	"example.com/blogArch/gateway/utils"
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

		resp := responses.StatusResponse{
			Status: "In Main Page!",
		}
		c.JSON(http.StatusOK, resp)
	}
}

func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In profile controller...")
		user, err := utils.ExtractTokenUser(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp := responses.ProfileResponse{
			UserID:  user,
			Entries: utils.GrabEntries(),
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
		// Extract JWT
		user, err := utils.ExtractTokenUser(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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

		log.Printf("Ouput is %s", res.GetOutput())

		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		var status string
		if res.GetOutput() == "POSITIVE\n" {
			status = utils.InsertEntry(body.Entry, user)
		} else {
			status = "Entry not inserted. Please refrain from toxic comments."
		}
		resp := responses.StatusResponse{
			Status: status,
		}
		c.JSON(http.StatusOK, resp)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("In login controller...")
		var body models.LoginModel
		c.BindJSON(&body)

		status, err := utils.TryLogin(body.Username, body.Password)

		// return JWT token
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
			return
		}
		token, _ := utils.GenerateJWT(body.Username)
		c.JSON(http.StatusOK, gin.H{"token":token, "Status": status})
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
