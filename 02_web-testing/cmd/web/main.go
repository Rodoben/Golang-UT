package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// setup am app config, app will become the reciever for sharing information
	app := application{}
	// get application routes
	mux := app.routes()

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
