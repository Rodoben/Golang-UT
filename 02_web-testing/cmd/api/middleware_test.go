package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-testing/pkg/data"
)

func TestEnableCors(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	tests := []struct {
		name           string
		method         string
		expectedHeader bool
	}{
		{"get", "GET", false},
		{"preflight", "OPTIONS", true},
	}

	for _, v := range tests {
		handlerToTest := app.enableCors(nextHandler)
		req, _ := http.NewRequest(v.method, "http://test", nil)
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)

		if v.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") == "" {
			t.Errorf("%s: expected Header, but did not find it", v.name)
		}
		if !v.expectedHeader && rr.Header().Get("Access-Control-Allow-Credentials") != "" {
			t.Errorf("%s: expected no header, but got one", v.name)
		}
	}

}

func TestAuthRequired(t *testing.T) {
	user := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "Admin@example.com",
	}

	token, _ := app.generateTokenPair(&user)

	nexthandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	tests := []struct {
		name             string
		token            string
		expectAuthorized bool
		setHeader        bool
	}{
		{name: "valid token", token: fmt.Sprintf("Bearer %s", token.Token), expectAuthorized: true, setHeader: true},
		{name: "no token", token: "", expectAuthorized: false, setHeader: false},
		{name: "invalid token", token: fmt.Sprintf("Bearer %s", expiredToken), expectAuthorized: false, setHeader: true},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "http://test", nil)
		if test.setHeader {
			req.Header.Set("Authorization", test.token)
		}
		rr := httptest.NewRecorder()
		handler := app.authRequired(nexthandler)
		handler.ServeHTTP(rr, req)
		fmt.Println("_", rr.Code)
		if test.expectAuthorized && rr.Code == http.StatusUnauthorized {
			t.Errorf("%s: got code 401, and should not have", test.name)
		}
		if !test.expectAuthorized && rr.Code != http.StatusUnauthorized {
			t.Errorf("%s: got code 401, and should not have", test.name)
		}
	}
}
