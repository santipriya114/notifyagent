package service

import "log"

func ReceiveMessage(msg string) {
	log.Println("[CONSUMER] Message Received:", msg)

	ProcessNotification(
		"SMS",
		"+917483935838",
		"Hello Monika",
	)
}