package service

import (
	"strings"
	"testing"
)

func TestSendEmailRequiresConfiguredProvider(t *testing.T) {
	t.Setenv("SENDGRID_API_KEY", "")
	t.Setenv("FROM_EMAIL", "")
	t.Setenv("TO_EMAIL", "")

	err := SendEmail("Hello")
	if err == nil || !strings.Contains(err.Error(), "SENDGRID_API_KEY") {
		t.Fatalf("expected missing SENDGRID_API_KEY error, got %v", err)
	}
}
