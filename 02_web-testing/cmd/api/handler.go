package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"web-testing/pkg/data"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Credentails struct {
	UserName string `json:"email"`
	Password string `json:"password"`
}

func (app *application) authenticate(w http.ResponseWriter, req *http.Request) {

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

func (app *application) refreshToken(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	refreshToken := req.Form.Get("refresh_token")
	claims := &Claims{}
	_, err = jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.JWTSecret), nil
	})
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if time.Unix(claims.ExpiresAt.Unix(), 0).Sub(time.Now()) > 30*time.Second {
		app.errorJSON(w, errors.New("refresh token does not need renewed yet"), http.StatusTooEarly)
		return
	}
	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.DB.GetUser(userId)
	if err != nil {
		app.errorJSON(w, errors.New("unknown user"), http.StatusBadRequest)
		return
	}
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "__Host-refresh_token",
		Path:     "/",
		Value:    tokenPairs.RefreshToken,
		Expires:  time.Now().Add(refreshTokenExpiry),
		MaxAge:   int(refreshTokenExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   true,
	})
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

func (app *application) allUser(w http.ResponseWriter, req *http.Request) {
	users, err := app.DB.AllUsers()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	app.writeJSON(w, http.StatusOK, users)
}
func (app *application) GetOneUser(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(req, "userID"))
	fmt.Println("1", userID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	fmt.Println("2", userID)
	//userId = 1
	user, err := app.DB.GetUser(userID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	fmt.Println("3")
	_ = app.writeJSON(w, http.StatusOK, user)

}

func (app *application) insertUser(w http.ResponseWriter, req *http.Request) {
	var user data.User
	err := app.readJSON(w, req, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	_, err = app.DB.InsertUser(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (app *application) deleteUser(w http.ResponseWriter, req *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(req, "userID"))
	if err != nil {
		app.readJSON(w, req, http.StatusBadRequest)
		return
	}
	err = app.DB.DeleteUser(userId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func (app *application) updateUser(w http.ResponseWriter, req *http.Request) {

	var user data.User
	err := app.readJSON(w, req, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.UpdateUser(user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
