package grpc

import (
	"context"
	"errors"
	"log"
	"os"

	"notification-agent/internal/service"
	pb "notification-agent/proto"

	"google.golang.org/grpc/metadata"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
}

func validateToken(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("missing metadata")
	}

	token := md.Get("authorization")

	if len(token) == 0 {
		return errors.New("missing token")
	}

	expectedToken := os.Getenv("GRPC_TOKEN")

	if token[0] != "Bearer "+expectedToken {
		return errors.New("unauthorized")
	}

	return nil
}

func (s *Server) SendNotification(
	ctx context.Context,
	req *pb.NotificationRequest,
) (*pb.NotificationResponse, error) {

	if err := validateToken(ctx); err != nil {
		return nil, err
	}

	log.Println("Channel:", req.Channel)
	log.Println("Message:", req.Message)

	return &pb.NotificationResponse{
		Status: "SUCCESS",
	}, nil
}

func (s *Server) GetNotificationStatus(
	ctx context.Context,
	req *pb.StatusRequest,
) (*pb.StatusResponse, error) {

	status, err := service.GetNotificationStatus(
		req.NotificationId,
	)

	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		Status: status,
	}, nil
}