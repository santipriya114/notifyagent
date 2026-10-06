package service

import "testing"

func IsEmailEnabled(enabled bool) bool {
	return enabled
}

func IsSMSEnabled(enabled bool) bool {
	return enabled
}

func IsInAppEnabled(enabled bool) bool {
	return enabled
}

func TestEmailPreferenceEnabled(t *testing.T) {

	result := IsEmailEnabled(true)

	if !result {
		t.Fatal("Expected email preference to be enabled")
	}
}

func TestEmailPreferenceDisabled(t *testing.T) {

	result := IsEmailEnabled(false)

	if result {
		t.Fatal("Expected email preference to be disabled")
	}
}

func TestSMSPreferenceEnabled(t *testing.T) {

	result := IsSMSEnabled(true)

	if !result {
		t.Fatal("Expected SMS preference to be enabled")
	}
}

func TestInAppPreferenceEnabled(t *testing.T) {

	result := IsInAppEnabled(true)

	if !result {
		t.Fatal("Expected InApp preference to be enabled")
	}
}