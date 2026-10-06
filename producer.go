package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/joho/godotenv"
)

func ProduceMessage() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connStr := os.Getenv("SERVICEBUS_CONNECTION")
	


	client, err := azservicebus.NewClientFromConnectionString(
		connStr,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	sender, err := client.NewSender(
		"notifications",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = sender.SendMessage(
		context.Background(),
		&azservicebus.Message{
			Body: []byte("Hello Monika"),
		},
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Message sent successfully")
}