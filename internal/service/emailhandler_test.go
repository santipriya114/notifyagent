package service

import (
	"strings"
	"testing"
)

func TestValidateSendGridConfig(t *testing.T) {
	t.Setenv("SENDGRID_API_KEY", "SG.test-key")
	t.Setenv("FROM_EMAIL", "sender@example.com")
	t.Setenv("TO_EMAIL", "recipient@example.com")
	if err := validateSendGridConfig(); err != nil {
		t.Fatalf("expected valid SendGrid config, got error: %v", err)
	}

	t.Setenv("FROM_EMAIL", "")
	err := validateSendGridConfig()
	if err == nil || !strings.Contains(err.Error(), "FROM_EMAIL") {
		t.Fatalf("expected missing FROM_EMAIL error, got %v", err)
	}
}
