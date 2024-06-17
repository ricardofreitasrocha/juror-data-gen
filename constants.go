package main

import "fmt"

type Constant string

const (
	jwtKey Constant = "[super-secret-key-bureau][super-secret-key-bureau][super-secret-key-bureau]"

	connectionString Constant = "postgresql://system:postgres@localhost:5432/juror?sslmode=disable"

	baseUrl Constant = "http://localhost:8080/api/v1/moj"

	healthCheckUrl        Constant = "http://localhost:8080/health"
	generatePoolNumberUrl Constant = baseUrl + "/pool-request/generate-pool-number"
	requestPoolUrl        Constant = baseUrl + "/pool-request/new-pool"
	summonVotersUrl       Constant = baseUrl + "/pool-create/create-pool"
	addResponseUrl        Constant = baseUrl + "/juror-paper-response/response"

	markAsRespondedUrl Constant = baseUrl + "/juror-paper-response/update-status/%s/CLOSED"
	disqualifyUrl      Constant = baseUrl + "/disqualify/juror/%s"
	markUndeliverable  Constant = baseUrl + "/undeliverable-response/%s"
	excusalUrl         Constant = baseUrl + "/excusal-response/juror/%s"
	deferralUrl        Constant = baseUrl + "/deferral-response/juror/%s"
)

func (c Constant) String(params ...interface{}) string {
	return fmt.Sprintf(string(c), params...)
}
