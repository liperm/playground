package main

import (
	"log"
	"net"

	"github.com/liperm/go-grpc-service/internal/handlers"
	"github.com/liperm/go-grpc-service/pb"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServer(s, &handlers.UserServer{})
	pb.RegisterItemServer(s, &handlers.ItemServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
