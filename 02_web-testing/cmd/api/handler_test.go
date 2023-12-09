package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_app_Authentication(t *testing.T) {
	testAuthhandler := []struct {
		name         string
		requestBody  string
		expectedCode int
	}{
		{name: "Valid User", requestBody: `{"email":"admin@example.com","password":"secret"}`, expectedCode: http.StatusOK},
		{name: "Not json", requestBody: `{"email":"admin@example.com","password":"secret}`, expectedCode: http.StatusUnauthorized},
		{name: "Empty json", requestBody: `{}`, expectedCode: http.StatusUnauthorized},
		{name: "empty email", requestBody: `{"email":""}`, expectedCode: http.StatusUnauthorized},
		{name: "empty password", requestBody: `{"email":"admin@example.com"}`, expectedCode: http.StatusUnauthorized},
		{name: "Invalid email", requestBody: `{"email":"admin1@example.com"}`, expectedCode: http.StatusUnauthorized},
	}

	for _, test := range testAuthhandler {
		reader := strings.NewReader(test.requestBody)
		req, _ := http.NewRequest(http.MethodPost, "/auth", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.authenticate)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedCode {
			t.Errorf("%s: authentication failed: expected status code: %d, but got: %d ", test.name, test.expectedCode, rr.Code)
		}
	}

}
