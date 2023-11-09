package handlers

import (
	"context"

	"github.com/liperm/go-grpc-service/pb"
)

type ItemServer struct {
	pb.UnimplementedItemServer
}

func (s *ItemServer) GetItem(ctx context.Context, request *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return &pb.GetItemResponse{Id: 2, Price: 2.5}, nil
}
