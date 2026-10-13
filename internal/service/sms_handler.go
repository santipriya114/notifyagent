package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func normalizePhoneNumber(raw string) string {
	value := strings.TrimSpace(raw)
	if value == "" {
		return ""
	}

	var digits strings.Builder
	for _, ch := range value {
		if ch == '+' {
			digits.WriteRune(ch)
			continue
		}
		if ch >= '0' && ch <= '9' {
			digits.WriteRune(ch)
		}
	}

	value = digits.String()
	if value == "" {
		return ""
	}

	if strings.HasPrefix(value, "00") {
		value = "+" + value[2:]
	}
	if !strings.HasPrefix(value, "+") {
		switch {
		case len(value) == 10:
			value = "+91" + value
		case len(value) == 11 && strings.HasPrefix(value, "1"):
			value = "+" + value
		case len(value) == 12 && strings.HasPrefix(value, "91"):
			value = "+" + value
		}
	}
	return value
}

func resolveSMSFromNumber() string {
	for _, key := range []string{"SMS_FROM", "TWILIO_PHONE_NUMBER"} {
		if value := normalizePhoneNumber(os.Getenv(key)); value != "" {
			return value
		}
	}
	return ""
}

func resolveSMSToNumber(input string) string {
	for _, value := range []string{input, os.Getenv("SMS_TO"), os.Getenv("TO_PHONE")} {
		if normalized := normalizePhoneNumber(value); normalized != "" {
			return normalized
		}
	}
	return ""
}

func validateTwilioConfig(to string) error {
	if strings.TrimSpace(to) == "" {
		return errors.New("destination phone number is missing; set SMS_TO or pass a valid number")
	}
	if strings.TrimSpace(os.Getenv("TWILIO_ACCOUNT_SID")) == "" {
		return errors.New("TWILIO_ACCOUNT_SID is not configured")
	}
	if strings.TrimSpace(os.Getenv("TWILIO_AUTH_TOKEN")) == "" {
		return errors.New("TWILIO_AUTH_TOKEN is not configured")
	}
	if strings.TrimSpace(resolveSMSFromNumber()) == "" {
		return errors.New("TWILIO_PHONE_NUMBER or SMS_FROM is not configured")
	}
	return nil
}

func SendSMS(to string, message string) error {
	to = resolveSMSToNumber(to)
	if err := validateTwilioConfig(to); err != nil {
		return err
	}

	fromNumber := resolveSMSFromNumber()
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{
			Username: os.Getenv("TWILIO_ACCOUNT_SID"),
			Password: os.Getenv("TWILIO_AUTH_TOKEN"),
		},
	)

	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(fromNumber)
	params.SetBody("sms_appointment_reminders")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		if strings.Contains(err.Error(), "572002") {
			return fmt.Errorf("Twilio trial account restriction: %w. Verify the destination number in Twilio or upgrade the account", err)
		}
		log.Println("TWILIO ERROR:", err)
		return err
	}

	log.Println("SMS Sent Successfully")
	if resp != nil && resp.Sid != nil {
		log.Println("SID:", *resp.Sid)
	}

	return nil
}
