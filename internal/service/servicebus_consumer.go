package service

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func ReceiveMessages() {

	connStr := os.Getenv("SERVICEBUS_CONNECTION")

	client, err := azservicebus.NewClientFromConnectionString(
		connStr,
		nil,
	)
	if err != nil {
		panic(err)
	}

	receiver, err := client.NewReceiverForQueue(
		"notifications",
		nil,
	)
	if err != nil {
		panic(err)
	}

	for {

		messages, err := receiver.ReceiveMessages(
			context.Background(),
			1,
			nil,
		)

		if err != nil {
			log.Println("Receive Error:", err)
			continue
		}

		for _, msg := range messages {

			message := string(msg.Body)

			// Azure Service Bus Message ID
			log.Println("Azure Message ID:", msg.MessageID)

			log.Println("Received From Queue:", message)

			ProcessNotification(
				"EMAIL",
				"",
				message,
			)

			err = receiver.CompleteMessage(
				context.Background(),
				msg,
				nil,
			)

			if err != nil {
				log.Println(
					"Failed to complete message:",
					err.Error(),
				)
			}
		}
	}
}