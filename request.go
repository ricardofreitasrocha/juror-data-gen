package main

import (
	"bytes"
	"errors"
	"net/http"
)

func request(method string, url string, payload []byte, needsToken bool) (*http.Response, error) {
	var token string

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if needsToken {
		token, _ = makeToken("", true)

		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if response.StatusCode > 299 {
		return nil, errors.New("Request failed with status code: " + response.Status)
	}

	return response, err
}
