package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var testHandlers = []struct {
		name                    string
		route                   string
		expectedStatusCode      int
		expectedUrl             string
		expectedFirstStatusCode int
	}{
		{name: "home", route: "/", expectedStatusCode: http.StatusOK, expectedUrl: "/", expectedFirstStatusCode: http.StatusOK},
		{name: "not-found", route: "/fish", expectedStatusCode: http.StatusNotFound, expectedUrl: "/fish", expectedFirstStatusCode: http.StatusNotFound},
		{name: "profile", route: "/user/profile", expectedStatusCode: http.StatusOK, expectedUrl: "/", expectedFirstStatusCode: http.StatusTemporaryRedirect},
	}

	mux := app.routes()

	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for _, test := range testHandlers {
		fmt.Println("_______", ts.URL+test.route)
		resp, err := ts.Client().Get(ts.URL + test.route)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
		fmt.Println("_______", resp.Body)
		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("for %s: expected status code %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
		}

		if resp.Request.URL.Path != test.expectedUrl {
			t.Errorf("for %s: expected final url %s, but got %s", test.name, test.expectedUrl, resp.Request.URL.Path)
		}

		resp2, _ := client.Get(ts.URL + test.route)
		if resp2.StatusCode != test.expectedFirstStatusCode {
			t.Errorf("%s: expected first returned status code to be %d but got %d", test.name, test.expectedFirstStatusCode, resp2.StatusCode)
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

func Test_UploaduserImage(t *testing.T) {
	// set up pipe
	pr, pw := io.Pipe()
	// create a new writer of type *io.Writer
	writer := multipart.NewWriter(pw)

	// create a waitgroup and add 1 to it
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// simulate uploading a file using a go routine and our writer
	go simulateImageUpload("./testdata/img.png", writer, t, wg)

	// read from the pipe which receives data
	request := httptest.NewRequest("POST", "/", pr)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	// call app.uploadedFiles
	uploadedFiles, err := app.UploadFiles(request, "./testdata/uploads/")
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].OriginalFileName)); os.IsNotExist(err) {
		t.Errorf("expected file to exist: %s", err.Error())
	}

	// clean up
	_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].OriginalFileName))
	wg.Wait()
}

func simulateImageUpload(fileToUpload string, writer *multipart.Writer, t *testing.T, wg *sync.WaitGroup) {
	defer writer.Close()
	defer wg.Done()

	// create the form data filed 'file' with value being filename
	part, err := writer.CreateFormFile("file", path.Base(fileToUpload))
	if err != nil {
		t.Error(err)
		return
	}

	// open the actual file
	f, err := os.Open(fileToUpload)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	// decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		t.Error("error decoding image:", err)
	}

	// write the png to our io.Writer
	err = png.Encode(part, img)
	if err != nil {
		t.Error(err)
	}

}
