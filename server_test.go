package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var connectionCount = 0

const login = `{"username":"test-username","password":"test-password"}`
const falseLogin = `{"username":"false","password":"none"}`

func init() {
	go main()
	waitForConnection()
}

func TestAuthenticate(t *testing.T) {
	// authenticate first with username and password
	req, err := http.NewRequest("POST", "http://127.0.0.1:8700/authenticate", bytes.NewBuffer([]byte(login)))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := httpClient().Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, res.StatusCode == http.StatusOK, res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	// unmarshal returned json
	var token JwtToken
	json.Unmarshal(body, &token)

	// get protected content with returned token
	req, err = http.NewRequest("GET", "http://127.0.0.1:8700/protected", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.Token)
	res, err = httpClient().Do(req)
	if err != nil {
		t.Error(err)
	}
	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()

	// assert that username and password are the same as first sent
	assert.Equal(t, login, strings.TrimSpace(string(body)), string(body))
}

func TestFalseAuthenticate(t *testing.T) {
	req, err := http.NewRequest("POST", "http://127.0.0.1:8700/authenticate", bytes.NewBuffer([]byte(falseLogin)))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := httpClient().Do(req)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, res.StatusCode == http.StatusOK, res.Status)
	res.Body.Close()
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(5*time.Second))
}

func httpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: dialTimeout,
		},
	}
}

func get(t *testing.T, url string) {
	resp, err := httpClient().Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	assert.True(t, resp.StatusCode == 200, "status code was not 200")
}

func waitForConnection() {
	if connectionCount < 10 {
		resp, err := http.Get("http://localhost:8700/health")
		if err != nil || resp.StatusCode != 200 {
			time.Sleep(1 * time.Second)
			connectionCount++
			waitForConnection()
		}
	} else {
		log.Println("Connecting to http://localhost:8700/health failed")
		os.Exit(1)
	}
}
