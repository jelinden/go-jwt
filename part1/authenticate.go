package main

import (
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

// username and password should be fetched from a database
const username = "test-username"
const password = "test-password"

// a secret from an environment variable for example, should not
// be in public version control
const secretKey = "secret"

func authenticate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	if validateCredentials(user) {
		json.NewEncoder(w).Encode(JwtToken{Token: signedTokenString(user)})
		return
	}
	json.NewEncoder(w).Encode(Exception{Message: "invalid credentials"})
}

func validateCredentials(user User) bool {
	if user.Username == username && user.Password == password {
		return true
	}
	return false
}

func signedTokenString(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
