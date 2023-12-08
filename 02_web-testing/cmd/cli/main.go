package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type application struct {
	JWTSecret string
	Action    string
}

// This is used to generate a token, so that we can test our api. Run this with go run ./cmd/cli and copy
// the token that is printed out.
// go run ./cmd/cli -action=valid     // will produce a valid token
// go run ./cmd/cli -action=expired   // will produce an expired token

func main() {
	var app application

	app.JWTSecret = `"jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "signing secret"`
	app.Action = `"action","valid","action: valid|expired"`

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "ronald benjamin"
	claims["sub"] = "1"
	claims["aud"] = "example.com"
	claims["iss"] = "example.com"
	claims["admin"] = true
	if app.Action == "valid" {
		expires := time.Now().UTC().Add(time.Hour * 72)
		claims["exp"] = expires.Unix()
	} else {
		expires := time.Now().UTC().Add(time.Hour * 72)
		claims["exp"] = expires.Unix()
	}
	// create the token as a slice of bytes
	if app.Action == "valid" {
		fmt.Println("VALID Token:")
	} else {
		fmt.Println("EXPIRED Token:")
	}

	// create a signed token
	signedAccessToken, err := token.SignedString([]byte(app.JWTSecret))
	if err != nil {
		log.Fatal(err)
	}
	// print to console
	fmt.Println(string(signedAccessToken))
}
