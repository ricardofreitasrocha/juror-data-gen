package main

import (
	"strings"
	"testing"
)

func TestMakeToken(t *testing.T) {
	token, _ := makeToken("415", true)

	if token == "" {
		t.Errorf("Expected token to be generated")
	}

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		t.Errorf("Expected token to have 3 parts")
	}
}

func TestMakeAdminToken(t *testing.T) {
	token, _ := makeToken("415", true, "ADMIN")

	if token == "" {
		t.Errorf("Expected token to be generated")
	}

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		t.Errorf("Expected token to have 3 parts")
	}
}

func TestBureauJwt(t *testing.T) {
	claims := Claims{}
	bureauJwt(&claims)

	if claims.Owner != "400" {
		t.Errorf("Expected owner to be 400")
	}

	if claims.Login != "bureau-lead" {
		t.Errorf("Expected login to be bureau-lead")
	}

	if claims.UserType != "BUREAU" {
		t.Errorf("Expected user type to be BUREAU")
	}

	if claims.ActiveUserType != "BUREAU" {
		t.Errorf("Expected active user type to be BUREAU")
	}

	if claims.Staff.Name != "bureau-lead" {
		t.Errorf("Expected staff name to be bureau-lead")
	}

	if claims.LocCode != "400" {
		t.Errorf("Expected loc code to be 400")
	}
}

func TestCourtJwt(t *testing.T) {
	claims := Claims{}
	courtJwt(&claims, "415")

	if claims.Owner != "415" {
		t.Errorf("Expected owner to be 415")
	}

	if claims.Login != "court-manager" {
		t.Errorf("Expected login to be court-manager")
	}

	if claims.UserType != "COURT" {
		t.Errorf("Expected user type to be COURT")
	}

	if claims.Staff.Name != "court-manager" {
		t.Errorf("Expected staff name to be court-manager")
	}

	if claims.LocCode != "415" {
		t.Errorf("Expected loc code to be 415")
	}
}

func TestGeneratePoolNumber(t *testing.T) {
	locCode := "415"

	poolNumber, _ := generatePoolNumber(locCode)

	if poolNumber == "" {
		t.Errorf("Expected pool number to be generated")
	}

	if len(poolNumber) != 9 {
		t.Errorf("Expected pool number to be 9 characters")
	}

	if poolNumber[:3] != locCode {
		t.Errorf("Expected pool number to start with location code")
	}
}
