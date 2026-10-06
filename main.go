package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	grpcserver "notification-agent/internal/grpc"
	"notification-agent/internal/handler"
	"notification-agent/internal/repository"
	"notification-agent/internal/service"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
}

func sendHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("[INFO] SEND API HIT")

	// Trigger notification flow
	service.ReceiveMessage("Hello Monika")

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "SUCCESS",
		"message": "Notification processed",
	})
}

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect PostgreSQL
	repository.ConnectDB()

	// Create tables
	repository.CreateTables()

	// Start gRPC Server
	go grpcserver.StartGRPCServer()

	// Health API
	http.HandleFunc("/health", healthHandler)

	// Notification APIs
	http.HandleFunc(
		"/notifications",
		handler.RateLimit(
			handler.GetNotificationsHandler,
		),
	)

	http.HandleFunc(
		"/create-notification",
		handler.RateLimit(
			handler.CreateNotification,
		),
	)

	http.HandleFunc(
		"/send",
		handler.RateLimit(
			sendHandler,
		),
	)

	log.Println("[INFO] HTTP Server started on :8085")

	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Println("[ERROR]", err.Error())
	}
}