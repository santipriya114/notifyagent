package service

import (
	"log"
	"os"
	"strings"
)

func Worker(
	id int,
	jobs chan Message,
	dlq chan Message,
) {
	for msg := range jobs {
		log.Printf("[WORKER %d] Processing: %s", id, msg.Content)

		smsTo := strings.TrimSpace(os.Getenv("SMS_TO"))
		if smsTo == "" {
			smsTo = strings.TrimSpace(os.Getenv("TO_PHONE"))
		}
		if smsTo == "" {
			log.Printf("[WARN] SMS recipient is not configured; set SMS_TO or TO_PHONE to a verified Twilio number")
		} else {
			err := SendSMS(smsTo, msg.Content)
			if err != nil {
				log.Printf("[ERROR] SMS failed: %v", err)
			} else {
				log.Printf("[SUCCESS] SMS sent")
			}
		}

		err := SendEmail(msg.Content)
		if err != nil {
			log.Printf("[ERROR] Email failed: %v", err)
		} else {
			log.Printf("[SUCCESS] Email sent")
		}
	}
}

func StartWorkerPool() chan Message {
	jobs := make(chan Message, 100)

	for i := 1; i <= 3; i++ {
		go Worker(
			i,
			jobs,
			DLQ,
		)
	}

	go ProcessDLQ()

	return jobs
}
