package service

import "testing"

func TestSendEmailSuccess(t *testing.T) {

	err := SendEmail("Hello")

	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}