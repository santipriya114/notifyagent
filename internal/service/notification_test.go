package service

import "testing"

func TestProcessNotificationEmail(t *testing.T) {

	ProcessNotification(
		"EMAIL",
		"",
		"Hello Monika",
	)
}

func TestProcessNotificationSMS(t *testing.T) {

	ProcessNotification(
		"SMS",
		"+917483935838",
		"Hello Monika",
	)
}

func TestProcessNotificationInApp(t *testing.T) {

	ProcessNotification(
		"INAPP",
		"",
		"Hello Monika",
	)
}

func TestProcessNotificationUnsupportedChannel(t *testing.T) {

	ProcessNotification(
		"UNKNOWN",
		"",
		"Hello Monika",
	)
}