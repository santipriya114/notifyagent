package service

import "testing"

func TestProcessNotificationHandlesMissingConfig(t *testing.T) {
	t.Setenv("SENDGRID_API_KEY", "")
	t.Setenv("FROM_EMAIL", "")
	t.Setenv("TO_EMAIL", "")
	t.Setenv("TWILIO_ACCOUNT_SID", "")
	t.Setenv("TWILIO_AUTH_TOKEN", "")
	t.Setenv("TWILIO_PHONE_NUMBER", "")

	ProcessNotification("EMAIL", "", "Hello Monika")
	ProcessNotification("SMS", "+14155552671", "Hello Monika")
	ProcessNotification("INAPP", "", "Hello Monika")
	ProcessNotification("UNKNOWN", "", "Hello Monika")
}
