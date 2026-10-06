package grpc

import (
	"log"
	"net"

	pb "notification-agent/proto"

	"google.golang.org/grpc"
)

func StartGRPCServer() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterNotificationServiceServer(
		server,
		&Server{},
	)

	log.Println("gRPC Server started on :50051")

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}