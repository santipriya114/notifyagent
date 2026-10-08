package service

import "testing"

func TestSendEmailSuccessHandler(t *testing.T) {

	err := SendEmail("Hello Test")

	if err != nil {
		t.Fatalf("SendEmail failed: %v", err)
	}
}