package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// setup am app config, app will become the reciever for sharing information
	app := application{
		Session: getSession(),
	}
	// get application routes
	mux := app.routes()
	// get a session manager
	//app.Session = getSession()
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
