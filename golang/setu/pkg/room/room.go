package room

import (
	"context"
	"fmt"
	pb "townhall-backend/api/v1"
)

type RoomService struct {
	pb.UnimplementedRoomServiceServer
	rooms map[string][]string
}

func NewRoomService() *RoomService {
	return &RoomService{
		rooms: make(map[string][]string),
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	roomID := fmt.Sprintf("%s_room", req.RoomName)
	s.rooms[roomID] = []string{}
	return &pb.CreateRoomResponse{
		RoomId: roomID,
	}, nil
}

func (s *RoomService) GetRoomInfo(ctx context.Context, req *pb.GetRoomInfoRequest) (*pb.GetRoomInfoResponse, error) {
	users, ok := s.rooms[req.RoomId]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}
	return &pb.GetRoomInfoResponse{
		Users: users,
	}, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, req *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	users, ok := s.rooms[req.RoomId]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}
	s.rooms[req.RoomId] = append(users, req.UserId)
	return &pb.JoinRoomResponse{
		RoomId: req.RoomId,
		Users:  s.rooms[req.RoomId],
	}, nil
}
