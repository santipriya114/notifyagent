package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "notification-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewNotificationServiceClient(conn)

	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		"authorization",
		"Bearer "+os.Getenv("GRPC_TOKEN"),
	)

	resp, err := client.SendNotification(
		ctx,
		&pb.NotificationRequest{
			Channel: "EMAIL",
			Message: "Hello gRPC",
		},
		grpc.WaitForReady(true),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", resp.Status)

	time.Sleep(time.Second)
}
