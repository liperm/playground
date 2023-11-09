package handlers

import (
	"context"

	"github.com/liperm/go-grpc-service/pb"
)

type UserServer struct {
	pb.UnimplementedUserServer
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUseRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Id: 2, Name: "claudio"}, nil
}
