package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var testHandlers = []struct {
		name               string
		route              string
		expectedStatusCode int
	}{
		{name: "sucess", route: "/", expectedStatusCode: http.StatusOK},
		{name: "not-found", route: "/fish", expectedStatusCode: http.StatusNotFound},
	}

	mux := app.routes()

	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	for _, test := range testHandlers {
		resp, err := ts.Client().Get(ts.URL + test.route)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s: expected status code %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
		}

	}

}

func Test_AppHome(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)
	handler.ServeHTTP(rr, req)
	//check status code
	if rr.Code != http.StatusOK {
		t.Errorf("TestAppHome returned wrong status code; expected 200 but got %d", rr.Code)
	}

	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), `<small>Your request came from Session:`) {
		t.Error("did not find the correct text in html")
	}
}

func Test_AppProfile(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Profile)
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("TestAppProfile returned wrong status code; expected 200 but got %d", rr.Code)
	}
	body, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(body), `<h1 class="mt-3">User Profile`) {
		t.Error("did not find the correct text in html")
	}
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))
	return req.WithContext(ctx)
}

func getCtx(req *http.Request) context.Context {
	return context.WithValue(req.Context(), contextUserKey, "test")
}

func Test_Table_AppHome(t *testing.T) {

	var tests = []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{name: "first visit", putInSession: "", expectedHTML: "<small>Your request came from Session:"},
		{name: "second visit", putInSession: "hello,world!", expectedHTML: "<small>Your request came from Session: hello,world!"},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req = addContextAndSessionToRequest(req, app)
		app.Session.Destroy(req.Context())

		if test.putInSession != "" {
			app.Session.Put(req.Context(), "test", test.putInSession)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Home)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("TestAppHome returned wrong status code; expected 200 but got %d", rr.Code)
		}

		body, _ := io.ReadAll(rr.Body)
		if !strings.Contains(string(body), test.expectedHTML) {
			t.Error("did not find the correct text in html")
		}
	}

}

func Test_App_renderWithBadTemplate(t *testing.T) {
	pathToTemplates = "./testData/"

	req, _ := http.NewRequest("GET", "/", nil)
	req = addContextAndSessionToRequest(req, app)
	rr := httptest.NewRecorder()
	err := app.render(rr, req, "bad.page.gohtml", &TemplateData{})
	if err == nil {
		t.Error("expected error from bad template, burt not got one")
	}
}

func Test_Login(t *testing.T) {
	testLogin := []struct {
		name               string
		postedData         Form
		expectedStatusCode int
		expectedLocation   string
	}{
		{
			name: "Missing form data",
			postedData: Form{Data: url.Values{
				"email":    {""},
				"password": {""},
			}},
			expectedStatusCode: 303,
			expectedLocation:   "/",
		},
		{
			name: "Bad credentials",
			postedData: Form{Data: url.Values{
				"email":    {"admin@admin.com"},
				"password": {"secret"},
			}},
			expectedStatusCode: 303,
			expectedLocation:   "/",
		},
		{
			name: "User not found",
			postedData: Form{Data: url.Values{
				"email":    {"admin2@example.com"},
				"password": {"secret"},
			}},
			expectedStatusCode: 303,
			expectedLocation:   "/",
		},
		{
			name: "authentication",
			postedData: Form{Data: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secrett"},
			}},
			expectedStatusCode: 303,
			expectedLocation:   "/",
		},
		{
			name: "valid login",
			postedData: Form{Data: url.Values{
				"email":    {"admin@example.com"},
				"password": {"secret"},
			}},
			expectedStatusCode: 303,
			expectedLocation:   "/user/profile",
		},
	}

	for _, test := range testLogin {
		req, _ := http.NewRequest("POST", "/login", strings.NewReader(test.postedData.Data.Encode()))
		req = addContextAndSessionToRequest(req, app)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Login)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatusCode {
			t.Errorf("%s: returned wrong status code; expected %d, but got %d", test.name, test.expectedStatusCode, rr.Code)
		}
		actualLocation, err := rr.Result().Location()
		if err == nil {
			if actualLocation.String() != test.expectedLocation {
				t.Errorf("%s: expected %s got %s", test.name, test.expectedLocation, actualLocation)
			}
		}
		t.Log("location:", actualLocation.String())

	}
}
