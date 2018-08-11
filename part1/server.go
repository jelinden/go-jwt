package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

func main() {
	router := httprouter.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	router.POST("/authenticate", authenticate)
	router.GET("/protected", authorizeMiddleware(http.HandlerFunc(protectedEndpoint)))
	router.GET("/health", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8700", router))
}

func protectedEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}
