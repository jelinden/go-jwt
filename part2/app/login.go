package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"html/template"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

var secretKey string
var credentials = make(map[string]string)
var tmpl = template.Must(template.ParseFiles("./src/index.html"))

func init() {
	secretKey = os.Getenv("JWT_KEY")
}

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tmpl.Execute(w, nil)
}

func LoginPOST(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	if validateCredentials(user) {
		json.NewEncoder(w).Encode(JwtToken{Token: signedTokenString(Username{Username: user.Username})})
		return
	}
	json.NewEncoder(w).Encode(Exception{Message: "invalid credentials"})
}

func validateCredentials(user User) bool {
	return credentials[user.Username] == user.Password
}

func signedTokenString(user Username) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
