package main

import (
	"fmt"
	"log"
	"net/http"
	"web-testing/pkg/repository"
	"web-testing/pkg/repository/dbrepo"
)

const (
	constantPort    = ":8090"
	host            = "localhost"
	user            = "postgres"
	password        = "postgres"
	dbname          = "users"
	sslmode         = "disable"
	timezone        = "UTC"
	connect_timeout = 5
	port            = 5432
)

type application struct {
	DSN string
	DB  repository.DatabaseRepo
}

func main() {
	var app application
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	app.DSN = dsn
	conn, err := app.connectToDatabase()
	if err != nil {
		log.Println("Unable to connect to databse")
	}
	defer conn.Close()
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	mux := app.Routes()
	log.Println("Api server running on port:", constantPort)
	err = http.ListenAndServe(constantPort, mux)
	if err != nil {
		panic(err)
	}

}
