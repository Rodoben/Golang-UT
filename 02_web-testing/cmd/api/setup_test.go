package main

import (
	"fmt"
	"os"
	"testing"
	"web-testing/pkg/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE3MDIzMDA5ODksImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6InJvbmFsZCBiZW5qYW1pbiIsInN1YiI6IjEifQ.deOPhFAMAh21Y7gOzI_KtHFkz1hXh5QHf3n-ki-kqRo"

func TestMain(m *testing.M) {
	app.DSN = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160"
	os.Exit(m.Run())
}
