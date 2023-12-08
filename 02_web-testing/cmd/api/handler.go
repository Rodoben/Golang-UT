package main

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentails struct {
	UserName string `json:"email"`
	Password string `json:"password"`
}

func (app *application) Authenticate(w http.ResponseWriter, req *http.Request) {

	// read json payload
	var data Credentails
	err := app.readJSON(w, req, &data)
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized1"), http.StatusUnauthorized)
		return
	}
	fmt.Println("1")
	// look up the user by email
	user, err := app.DB.GetUserByEmail(data.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized2"), http.StatusUnauthorized)
		return
	}
	fmt.Println("2")
	// checkpassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized3"), http.StatusUnauthorized)
		return
	}
	//generate tokens
	fmt.Println("3")
	tokenpairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized4"), http.StatusUnauthorized)
		return
	}
	fmt.Println("4")
	//send token to user
	_ = app.writeJSON(w, http.StatusOK, tokenpairs)
}
