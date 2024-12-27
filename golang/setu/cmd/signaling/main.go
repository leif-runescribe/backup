package main

import (
	"log"
	"net"
	pb "townhall-backend/api/v1"
	"townhall-backend/pkg/signaling"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSignalingServiceServer(grpcServer, &signaling.SignalingService{})
	log.Printf("Signaling server started on :50051")
	grpcServer.Serve(lis)
}
