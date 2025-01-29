package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	resRecorder := httptest.NewRecorder()

	handler(resRecorder, req)

	if status := resRecorder.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, instead of %v", status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(resRecorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Could not parse response: %v", err)
	}

	emailAddress := "oguejioforrichard@gmail.com"
	gitHubURL := "https://github.com/ikennarichard/hng12/stage-0"

	if response.Email != emailAddress {
		t.Errorf("Wrong email: got %v, instead of %v", response.Email, emailAddress)
	}

	if response.GitHubURL != gitHubURL {
		t.Errorf("Wrong github url: got %v, instead of %v", response.GitHubURL, gitHubURL)
	}

	_, err = time.Parse(time.RFC3339, response.CurrentDateTime)
	if err != nil {
		t.Error("Invalid CurrentDateTime format, expected ISO 8601 format", response.CurrentDateTime)
	}
}
