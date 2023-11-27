package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	registedRoutes := []struct {
		name   string
		route  string
		method string
	}{
		{name: "Route 1", route: "/", method: "GET"},
		{name: "Route 2", route: "/login", method: "POST"},
		{name: "Route 3", route: "/user/profile", method: "GET"},
		{name: "Static Route", route: "/static/*", method: "GET"},
	}

	mux := app.routes()
	chiRoutes := mux.(chi.Routes)

	for _, route := range registedRoutes {
		if !routeExist(route.route, route.method, chiRoutes) {
			t.Errorf("%s route %s is not registered", route.name, route.route)
		}
	}
}

func routeExist(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false
	_ = chi.Walk(chiRoutes, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})
	return found
}
