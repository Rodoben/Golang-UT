package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-testing/pkg/data"
)

func Test_GetTokenFromHeaderAndVerify(t *testing.T) {

	user := data.User{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@example.com",
		Password:  "secret",
		IsAdmin:   1,
	}

	tokens, _ := app.generateTokenPair(&user)

	tests := []struct {
		name          string
		token         string
		errorExpected bool
		setHeader     bool
		issuer        string
	}{
		{name: "Valid Token", token: fmt.Sprintf("Bearer %s", tokens.Token), errorExpected: false, setHeader: true, issuer: app.Domain},
		{name: "Valid Expired", token: fmt.Sprintf("Bearer %s", expiredToken), errorExpected: true, setHeader: true, issuer: app.Domain},
		{name: "invalid Token", token: fmt.Sprintf("Bearer %s1", tokens.Token), errorExpected: true, setHeader: true, issuer: app.Domain},
		{name: "no Bearer", token: fmt.Sprintf("Beare %s", tokens.Token), errorExpected: true, setHeader: true, issuer: app.Domain},
		{name: "Three Parts", token: fmt.Sprintf("Bearer %s 1", tokens.Token), errorExpected: true, setHeader: true, issuer: app.Domain},
		{name: "wrong issuer", token: fmt.Sprintf("Bearer %s", tokens.Token), errorExpected: true, setHeader: true, issuer: "anotherdomain.com"},
		{name: "No Header", token: "", errorExpected: true, setHeader: false, issuer: app.Domain},
	}

	for _, test := range tests {
		if test.issuer != app.Domain {
			app.Domain = test.issuer
			tokens, _ = app.generateTokenPair(&user)
		}
		req, _ := http.NewRequest("POST", "test", nil)
		if test.setHeader {
			req.Header.Set("Authorization", test.token)
		}
		rr := httptest.NewRecorder()
		_, _, err := app.getTokenFromHeaderAndVerify(rr, req)
		if err != nil && !test.errorExpected {
			t.Errorf("%s: did not expect an error, but got one %s ", test.name, err.Error())
		}
		if err == nil && test.errorExpected {
			t.Errorf("%s: expected error, but did not get one", test.name)
		}
		app.Domain = "example.com"
	}
}
