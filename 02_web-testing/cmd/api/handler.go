package main

import (
	"errors"
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
	// look up the user by email
	user, err := app.DB.GetUserByEmail(data.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized2"), http.StatusUnauthorized)
		return
	}
	// checkpassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized3"), http.StatusUnauthorized)
		return
	}
	//generate tokens

	tokenpairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("Unauthorized4"), http.StatusUnauthorized)
		return
	}

	//send token to user
	_ = app.writeJSON(w, http.StatusOK, tokenpairs)
}
