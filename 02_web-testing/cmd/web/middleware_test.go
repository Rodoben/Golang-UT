package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web-testing/pkg/data"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	nexthandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		t.Log(ip)
	})

	for _, e := range tests {
		handlerToTest := app.addIPToContext(nexthandler)
		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}
		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}
		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}

}

func Test_ipFromContext(t *testing.T) {

	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "somekey")
	ip := app.ipFromContext(ctx)

	if !strings.EqualFold(ip, "somekey") {
		t.Error("user-ip mismatch")
	}
}

func Test_middleware_Auth(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	var testAuth = []struct {
		name   string
		isAuth bool
	}{
		{name: "logged in", isAuth: true},
		{name: "not logged in", isAuth: false},
	}

	for _, test := range testAuth {
		handlerToTest := app.auth(nextHandler)
		// need a request
		req, _ := http.NewRequest("GET", "http://testing", nil)
		req = addContextAndSessionToRequest(req, app)
		if test.isAuth {
			app.Session.Put(req.Context(), "user", data.User{ID: 1})
		}
		rr := httptest.NewRecorder()
		handlerToTest.ServeHTTP(rr, req)
		if test.isAuth && rr.Code != http.StatusOK {
			t.Errorf("%s: expected status code of 200 but got %d", test.name, rr.Code)
		}
		if !test.isAuth && rr.Code != http.StatusTemporaryRedirect {
			t.Errorf("%s: expected status code of 307 but got %d", test.name, rr.Code)
		}
	}
}
