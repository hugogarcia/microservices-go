package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func Test_Authenticate(t *testing.T) {
	postBody := map[string]string{
		"email":    "em@example.com",
		"password": "XXXXXXXX",
	}

	jsonReturnMock := `{
		"error": false,
		"message": "some message"
	}`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(jsonReturnMock)),
		}
	})

	testApp.Client = client

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/auth", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.Authenticate)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d but got %d", http.StatusAccepted, rr.Code)
	}
}
