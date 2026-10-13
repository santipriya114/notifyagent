package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func validateSendGridConfig() error {
	if strings.TrimSpace(os.Getenv("SENDGRID_API_KEY")) == "" {
		return errors.New("SENDGRID_API_KEY is not configured")
	}
	if strings.TrimSpace(os.Getenv("FROM_EMAIL")) == "" {
		return errors.New("FROM_EMAIL is not configured; verify the sender identity in SendGrid")
	}
	if strings.TrimSpace(os.Getenv("TO_EMAIL")) == "" {
		return errors.New("TO_EMAIL is not configured")
	}
	return nil
}

func SendEmail(message string) error {
	if err := validateSendGridConfig(); err != nil {
		return err
	}

	fromName := strings.TrimSpace(os.Getenv("FROM_NAME"))
	if fromName == "" {
		fromName = "Notification Agent"
	}

	fromEmail := strings.TrimSpace(os.Getenv("FROM_EMAIL"))
	toEmail := strings.TrimSpace(os.Getenv("TO_EMAIL"))

	from := mail.NewEmail(fromName, fromEmail)
	to := mail.NewEmail("Recipient", toEmail)

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
		body := strings.TrimSpace(response.Body)
		if strings.Contains(body, "verified Sender Identity") || strings.Contains(body, "from address") {
			return fmt.Errorf("SendGrid rejected the request: %s. Verify %s as a Sender Identity in SendGrid before sending email", body, fromEmail)
		}
		log.Printf("[ERROR] SendGrid rejected the request")
		return fmt.Errorf("sendgrid error: %s", body)
	}

	log.Printf("[SUCCESS] Email sent successfully")
	return nil
}
