package main

import (
	"fmt"
	"os"
	"testing"

	"web-testing/pkg/repository/dbrepo"
)

// this file runs before other _test.go files.
var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"
	app.Session = getSession()
	app.DSN = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)

	app.DB = &dbrepo.TestDBRepo{}
	os.Exit(m.Run())
}
