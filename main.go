package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	grpcserver "notification-agent/internal/grpc"
	"notification-agent/internal/handler"
	"notification-agent/internal/service"
)

var jobs chan service.Message

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] SEND API HIT")

	message := r.URL.Query().Get("message")

	if message == "" {
		message = "Hello shantipriya"
	}

	jobs <- service.Message{
		Content: message,
		Retries: 0,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "SUCCESS",
		"message": message,
	})
}

func main() {

	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// PostgreSQL temporarily disabled
	// repository.ConnectDB()
	// repository.CreateTables()

	// Start Worker Pool
	jobs = service.StartWorkerPool()

	log.Println("[INFO] Worker Pool Started")

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
