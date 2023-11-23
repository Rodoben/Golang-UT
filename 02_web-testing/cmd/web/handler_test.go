package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
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
