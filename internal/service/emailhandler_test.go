package service

import "testing"

func TestSendEmailSuccessHandler(t *testing.T) {

	err := SendEmail("Hello Test")

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
