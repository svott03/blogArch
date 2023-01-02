package main

import (
	"context"
	"log"
	"net"
	"os/exec"

	pb "example.com/blogArch/proto"
	"google.golang.org/grpc"
)

const (
	// Port for gRPC server to listen to
	PORT = ":50051"
)

type TextFilterServer struct {
	pb.UnimplementedTextFilterServiceServer
}

func (s *TextFilterServer) CreateFilterOutput(ctx context.Context, in *pb.FilterInput) (*pb.FilterOutput, error) {
	log.Printf("Received: %v", in.GetInput())
	// Run model

	cmd := exec.Command("zsh", "-c", "python3 analysis.py " + in.GetInput())
	res, err := cmd.Output()
	res_str := string(res)
	if err != nil {
		log.Fatal(err.Error())
		res_str = "Error in Prediction"
	}
	
	out := &pb.FilterOutput{
		Output:        res_str,
	}

	return out, nil

}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterTextFilterServiceServer(s, &TextFilterServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}