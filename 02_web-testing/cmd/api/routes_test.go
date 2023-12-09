package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_route(t *testing.T) {
	testRoutes := []struct {
		name   string
		method string
		route  string
	}{
		{name: "Auth route", method: http.MethodPost, route: "/authentication"},
		{name: "refresh route", method: http.MethodPost, route: "/refresh-token"},
		{name: "AllUser route", method: http.MethodGet, route: "/users/"},
		{name: "GetUser route", method: http.MethodGet, route: "/users/{userID}"},
		{name: "Delete route", method: http.MethodDelete, route: "/users/{userID}"},
		{name: "Put route", method: http.MethodPut, route: "/users/"},
		{name: "Patch route", method: http.MethodPatch, route: "/users/"},
	}

	mux := app.Routes()
	chiMux := mux.(chi.Router)
	for _, test := range testRoutes {
		if !routeWalk(test.route, test.method, chiMux) {
			t.Errorf("%s: not found with method:%s and route:%s", test.name, test.method, test.route)
		}
	}
}

func routeWalk(testroute string, testmethod string, chiRoutes chi.Router) bool {
	found := false
	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(testmethod, method) && strings.EqualFold(testroute, route) {
			found = true
		}
		return nil
	})
	return found
}
