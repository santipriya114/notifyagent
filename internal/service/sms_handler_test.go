package service

import (
	"strings"
	"testing"
)

func TestValidateTwilioConfig(t *testing.T) {
	t.Setenv("TWILIO_ACCOUNT_SID", "AC123")
	t.Setenv("TWILIO_AUTH_TOKEN", "token123")
	t.Setenv("TWILIO_PHONE_NUMBER", "+1234567890")
	if err := validateTwilioConfig("+14155552671"); err != nil {
		t.Fatalf("expected valid Twilio config, got error: %v", err)
	}

	t.Setenv("TWILIO_PHONE_NUMBER", "")
	err := validateTwilioConfig("+14155552671")
	if err == nil || !strings.Contains(err.Error(), "TWILIO_PHONE_NUMBER") {
		t.Fatalf("expected missing TWILIO_PHONE_NUMBER error, got %v", err)
	}
}

func TestResolveSMSNumbers(t *testing.T) {
	t.Setenv("SMS_TO", "+919691129372")
	t.Setenv("SMS_FROM", "+17372212163")

	if got := resolveSMSToNumber(""); got != "+919691129372" {
		t.Fatalf("expected SMS_TO to be used, got %q", got)
	}
	if got := resolveSMSFromNumber(); got != "+17372212163" {
		t.Fatalf("expected SMS_FROM to be used, got %q", got)
	}
}

func TestNormalizePhoneNumber(t *testing.T) {
	if got := normalizePhoneNumber("9691129372"); got != "+919691129372" {
		t.Fatalf("expected +919691129372, got %q", got)
	}
	if got := normalizePhoneNumber("+17372212163"); got != "+17372212163" {
		t.Fatalf("expected +17372212163, got %q", got)
	}
}
