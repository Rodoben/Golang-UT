package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"web-testing/pkg/db"
)

// this file runs before other _test.go files.
var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	pathToTemplates = "./../../cmd/templates/"
	app.DSN = fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	conn, err := app.connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
