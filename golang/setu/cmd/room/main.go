package main

import (
	"log"
	"net"
	pb "townhall-backend/api/v1"
	"townhall-backend/pkg/room"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	roomService := room.NewRoomService()
	pb.RegisterRoomServiceServer(grpcServer, roomService)
	log.Printf("Room service started on :50052")
	grpcServer.Serve(lis)
}
