package main

import (
	"log"
	"net"

	grpcserver "notification-agent/internal/grpc"
	pb "notification-agent/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterNotificationServiceServer(
		s,
		&grpcserver.Server{},
	)

	log.Println("gRPC Server Started on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}