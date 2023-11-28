package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"web-testing/pkg/data"
	"web-testing/pkg/repository"
	"web-testing/pkg/repository/dbrepo"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
	DB      repository.DatabaseRepo
	DSN     string
}

const (
	host            = "localhost"
	user            = "postgres"
	password        = "postgres"
	dbname          = "users"
	sslmode         = "disable"
	timezone        = "UTC"
	connect_timeout = 5
	port            = 5432
)

func main() {
	gob.Register(data.User{})
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%v", host, port, user, password, dbname, sslmode, timezone, connect_timeout)
	//	setup an app config, app will become the reciever for sharing information
	app := application{
		Session: getSession(),
		DSN:     dsn,
	}
	// flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	// flag.Parse()
	conn, err := app.connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	// get application routes
	mux := app.routes()
	// get a session manager
	// app.Session = getSession()
	log.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
