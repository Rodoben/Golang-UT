package main

import (
	"fmt"
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.DSN = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	os.Exit(m.Run())
}
