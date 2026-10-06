package main

import (
	"context"
	"log"

	pb "notification-agent/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CheckStatus() {

	conn, err := grpc.Dial(
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

	resp, err := client.GetNotificationStatus(
		context.Background(),
		&pb.StatusRequest{
			NotificationId: 1,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Notification Status:", resp.Status)
}