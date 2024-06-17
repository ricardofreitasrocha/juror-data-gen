package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.Claims
	Owner           string   `json:"owner"`
	Login           string   `json:"login"`
	UserLevel       string   `json:"userLevel"`
	DaysToExpire    int      `json:"daysToExpire"`
	PasswordWarning bool     `json:"passwordWarning"`
	UserType        string   `json:"userType"`
	ActiveUserType  string   `json:"activeUserType"`
	Roles           []string `json:"roles"`
	Staff           struct {
		Name   string   `json:"name"`
		Rank   int      `json:"rank"`
		Active int      `json:"active"`
		Courts []string `json:"courts"`
	} `json:"staff"`
	LocCode string `json:"locCode"`
}

func makeToken(locCode string, bureau bool) (string, error) {
	claims := Claims{
		UserLevel:       "1",
		DaysToExpire:    999,
		PasswordWarning: false,
		Roles:           []string{"MANAGER"},
	}

	if bureau {
		bureauJwt(&claims)
	} else {
		courtJwt(&claims, locCode)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey.String()))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func bureauJwt(claims *Claims) {
	claims.Owner = "400"
	claims.Login = "bureau-lead"
	claims.UserType = "BUREAU"
	claims.ActiveUserType = "BUREAU"
	claims.Staff.Name = "bureau-lead"
	claims.LocCode = "400"
}

func courtJwt(claims *Claims, locCode string) {
	claims.Owner = locCode
	claims.Login = "court-manager"
	claims.UserType = "COURT"
	claims.ActiveUserType = "COURT"
	claims.Staff.Name = "court-manager"
	claims.LocCode = locCode
}

func generatePoolNumber(locCode string) (string, error) {
	queryParams := fmt.Sprintf("?locationCode=%s&attendanceDate=%s", locCode, "2024-05-08")

	result, err := request("GET", generatePoolNumberUrl.String()+queryParams, nil, true)
	if err != nil {
		return "", err
	}

	defer result.Body.Close()

	pnBuf := make([]byte, 9) // pool number is 9 characters long so 9 bytes is enough
	_, err = result.Body.Read(pnBuf)

	if err != nil && err.Error() != "EOF" {
		log.Error(err.Error())
		return "", err
	}

	poolNumber := string(pnBuf)

	return poolNumber, nil
}
