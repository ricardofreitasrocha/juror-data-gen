package main

import (
	"testing"
)

func TestMarkAsRespondedUrl(t *testing.T) {
	markAsRespondedUrl := markAsRespondedUrl.String("123456789")
	expectedUrl := "http://localhost:8080/api/v1/moj/juror-paper-response/update-status/123456789/CLOSED"

	print(markAsRespondedUrl)

	if markAsRespondedUrl != expectedUrl {
		t.Errorf("Expected markAsRespondedUrl to be %s", expectedUrl)
	}
}

func TestConstantWithNoFormatting(t *testing.T) {
	jwtKey := jwtKey.String()
	expected := "[super-secret-key-bureau][super-secret-key-bureau][super-secret-key-bureau]"

	if jwtKey != expected {
		t.Errorf("Expected jwtKet to be %s", expected)
	}
}
