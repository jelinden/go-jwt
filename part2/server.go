package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/jelinden/go-jwt/part2/app"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

func main() {
	router := httprouter.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	router.POST("/api/signup", app.SignupPOST)
	router.POST("/api/login", app.LoginPOST)
	router.GET("/api/profile", app.AuthorizeMiddleware(http.HandlerFunc(protectedEndpoint)))

	router.GET("/", app.Index)
	router.GET("/login", app.Index)
	router.GET("/signup", app.Index)

	router.GET("/health", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	router.Handler("GET", "/static/*filepath", http.StripPrefix("/static", http.FileServer(http.Dir("bin"))))

	log.Fatal(http.ListenAndServe(":8700", router))
}

func protectedEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user app.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}
