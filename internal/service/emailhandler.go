package service

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(message string) error {

	from := mail.NewEmail(
		"Monika",
		"monika.a.b@capgemini.com",
	)

	to := mail.NewEmail(
		"Monika",
		"temp1@yopmail.com",
	)

	subject := "Notification Alert"

	msg := mail.NewSingleEmail(
		from,
		subject,
		to,
		message,
		message,
	)

	client := sendgrid.NewSendClient(
		os.Getenv("SENDGRID_API_KEY"),
	)

	response, err := client.Send(msg)
	if err != nil {
		log.Printf("[ERROR] Email failed: %v", err)
		return err
	}

	log.Printf("Status Code: %d", response.StatusCode)
	log.Printf("Response Body: %s", response.Body)

	if response.StatusCode >= 400 {
		log.Printf("[ERROR] SendGrid rejected the request")
		return fmt.Errorf("sendgrid error: %s", response.Body)
	}

	log.Printf("[SUCCESS] Email sent successfully")
	return nil
}
