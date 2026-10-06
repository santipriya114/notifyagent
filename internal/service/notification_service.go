package service

import "notification-agent/internal/repository"

func CreateNotification(userID string, message string) (int64, error) {

	return repository.CreateNotification(
		userID,
		message,
	)
}

func ProcessNotification(
	channel string,
	phoneNumber string,
	message string,
) {

	switch channel {

	case "EMAIL":
		err := SendEmail(message)
		if err != nil {
			println("Email failed:", err.Error())
		}

	case "SMS":
		err := SendSMS(
			phoneNumber,
			message,
		)
		if err != nil {
			println("SMS failed:", err.Error())
		}

	case "INAPP":
		SendInApp(message)

	default:
		println("Unsupported channel")
	}
}