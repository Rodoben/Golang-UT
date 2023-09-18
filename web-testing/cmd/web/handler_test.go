package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},

		{"404-not-found", "/fish", http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	pathToTemplates = "./../../cmd/templates/"
	for _, e := range tests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s: expected status %d, but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
	}

}

func TestLogin(t *testing.T) {
	testCases := []struct {
		name     string
		request  *http.Request
		expected int
		body     string
	}{
		{
			name:     "Valid Request",
			request:  httptest.NewRequest("POST", "/login", strings.NewReader("email=test@example.com&password=12345")),
			expected: http.StatusOK,
			body:     "test@example.com",
		},
		{
			name:     "Missing Email",
			request:  httptest.NewRequest("POST", "/login", strings.NewReader("password=12345")),
			expected: http.StatusOK,
			body:     "failed validation",
		},
		{
			name:     "Missing Password",
			request:  httptest.NewRequest("POST", "/login", strings.NewReader("email=test@example.com")),
			expected: http.StatusOK,
			body:     "failed validation",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tc.request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			app := &application{} // Create your application instance here.

			app.Login(w, tc.request)
			if w.Code != tc.expected {
				t.Errorf("expected status code %d, got %d", tc.expected, w.Code)
			}

			if w.Body.String() != tc.body {
				t.Errorf("expected response body %s, got %s", tc.body, w.Body.String())
			}
		})
	}
}

func TestLogin1(t *testing.T) {
	// Create a request with form data.
	requestBody := "email=test@example.com&password=secretpassword"
	req := httptest.NewRequest("POST", "/login", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	// Create an instance of your application.
	app := &application{} // Replace with the actual initialization.

	// Call the Login method.
	app.Login(w, req)

	// Check the response.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponseBody := "test@example.com"
	if w.Body.String() != expectedResponseBody {
		t.Errorf("Expected response body %s, but got %s", expectedResponseBody, w.Body.String())
	}
}