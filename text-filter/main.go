package main

import (
	"context"
	"log"
	"net"

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

// TODO Add filter model and logic
func (s *TextFilterServer) CreateFilterOutput(ctx context.Context, in *pb.FilterInput) (*pb.FilterOutput, error) {
	log.Printf("Received: %v", in.GetInput())
	out := &pb.FilterOutput{
		Output:        in.GetInput() + "!",
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