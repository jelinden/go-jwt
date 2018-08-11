package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SignupPOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(err, "Oops, signup failure", w)
		return
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		handleError(err, "Oops, signup failure", w)
		return
	}
	if credentials[user.Username] != "" {
		handleError(errors.New("username exists"), "Username already exists", w)
		return
	}
	credentials[user.Username] = user.Password
	w.WriteHeader(200)
	w.Write([]byte(`{"status": "OK"}`))
}

func handleError(err error, msg string, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(200)
	e := User{Error: msg}
	s, _ := json.Marshal(e)
	w.Write(s)
}
