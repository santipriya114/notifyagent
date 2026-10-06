package service

import (
	"errors"
	"log"
	"time"
)

const MaxRetries = 3

func SendEmail(message string) error {

	var err error

	for i := 1; i <= MaxRetries; i++ {

		log.Printf("[INFO] Sending email... Attempt %d", i)

		if message == "" {
			err = errors.New("message is empty")
		} else {
			log.Printf("[INFO] Email sent successfully: %s", message)
			return nil
		}

		log.Printf("[ERROR] Attempt %d failed: %v", i, err)

		if i < MaxRetries {
			log.Println("[INFO] Retrying in 2 seconds...")
			time.Sleep(2 * time.Second)
		}
	}

	log.Printf("[ERROR] Email failed after %d attempts", MaxRetries)

	return err
}