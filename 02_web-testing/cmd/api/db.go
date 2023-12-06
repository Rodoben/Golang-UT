package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("error opening the fb connection: %s", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) connectToDatabase() (*sql.DB, error) {
	conn, err := openDB(app.DSN)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database!")
	return conn, nil
}
