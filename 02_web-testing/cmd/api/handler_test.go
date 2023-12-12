package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"web-testing/pkg/data"

	"github.com/go-chi/chi/v5"
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

func Test_app_RefreshToken(t *testing.T) {
	var tests = []struct {
		name               string
		token              string
		expectedCodeStatus int
		resetRefreshTime   bool
	}{
		{name: "valid", token: "", expectedCodeStatus: http.StatusOK, resetRefreshTime: true},
		{name: "expired token", token: expiredToken, expectedCodeStatus: http.StatusBadRequest, resetRefreshTime: false},
		{name: "valid but about to expire", token: "", expectedCodeStatus: http.StatusTooEarly, resetRefreshTime: false},
	}

	testUser := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
	}
	oldRefreshTime := refreshTokenExpiry
	for _, test := range tests {
		var tkn string
		if test.token == "" {
			if test.resetRefreshTime {
				refreshTokenExpiry = time.Second * 1
			}
			tokens, _ := app.generateTokenPair(&testUser)
			tkn = tokens.RefreshToken
		} else {
			tkn = test.token
		}

		postedData := url.Values{"refresh_token": {tkn}}

		req, _ := http.NewRequest("POST", "/refresh-token", strings.NewReader(postedData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()

		handlertoTest := http.HandlerFunc(app.refreshToken)
		handlertoTest.ServeHTTP(rr, req)

		if rr.Code != test.expectedCodeStatus {
			t.Errorf("%s: expected status code %d, but got %d", test.name, test.expectedCodeStatus, rr.Code)
		}
		refreshTokenExpiry = oldRefreshTime
	}
}

func Test_app_Databasehandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		json         string
		paramID      string
		handler      http.HandlerFunc
		expectedCode int
	}{
		{
			name:         "Insert user",
			method:       http.MethodPut,
			json:         `{"first_name":"Jack","last_name":"Smith","email":"jack@example.com"}`,
			handler:      app.insertUser,
			paramID:      "",
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Insert user",
			method:       http.MethodPut,
			json:         `{"first_name:"Jack","last_name":"Smith","email":"jack@example.com"}`,
			handler:      app.insertUser,
			paramID:      "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get User",
			method:       http.MethodGet,
			json:         "",
			paramID:      "1",
			handler:      app.GetOneUser,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Get User Invalid",
			method:       http.MethodGet,
			json:         "",
			paramID:      "2",
			handler:      app.GetOneUser,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get User Invalid",
			method:       http.MethodGet,
			json:         "",
			paramID:      "@",
			handler:      app.GetOneUser,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Get AllUser",
			method:       http.MethodGet,
			json:         "",
			paramID:      "",
			handler:      app.allUser,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Update User",
			method:       http.MethodPatch,
			json:         `{"id":1,"first_name":"Administrator","last_name":"User","email":"admin@example.com"}`,
			paramID:      "",
			handler:      app.updateUser,
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Update User invalid json",
			method:       http.MethodPatch,
			json:         `{"id:2,"first_name":"Administrator","last_name":"User","email":"admin@example.com"}`,
			paramID:      "",
			handler:      app.updateUser,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Update User invalid",
			method:       http.MethodPatch,
			json:         `{"id":2,"first_name":"Administrator","last_name":"User","email":"admin@example.com"}`,
			paramID:      "",
			handler:      app.updateUser,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Delete User",
			method:       http.MethodDelete,
			json:         "",
			paramID:      "1",
			handler:      app.deleteUser,
			expectedCode: http.StatusNoContent,
		},
	}

	for _, test := range tests {
		var req *http.Request
		if test.json == "" {
			req, _ = http.NewRequest(test.method, "/", nil)
		} else {
			req, _ = http.NewRequest(test.method, "/", strings.NewReader(test.json))
		}

		if test.paramID != "" {
			chiCtx := chi.NewRouteContext()
			chiCtx.URLParams.Add("userID", test.paramID)
			fmt.Println("----", test.paramID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
			fmt.Println("----", chiCtx.URLParams.Values, chiCtx.URLParams.Keys)
		}
		rr := httptest.NewRecorder()

		handlertoTest := http.HandlerFunc(test.handler)
		fmt.Println(handlertoTest)
		handlertoTest.ServeHTTP(rr, req)

		if rr.Code != test.expectedCode {
			t.Errorf("%s: handler expected code %d:, but got: %d", test.name, test.expectedCode, rr.Code)
		}
	}
}
