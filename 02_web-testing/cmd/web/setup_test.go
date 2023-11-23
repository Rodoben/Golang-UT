package main

import (
	"os"
	"testing"
)

// this file runs before other _test.go files.
var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	pathToTemplates = "./../../cmd/templates/"
	os.Exit(m.Run())
}
