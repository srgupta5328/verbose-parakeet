package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitialize(t *testing.T) {
	testRouter := Initialize()
	if testRouter == nil {
		t.Errorf("Error: '%v', Router didn't initialize", testRouter)
	}
}

func TestHomeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	HomeHandler(res, req)

	got := res.Body.String()
	want := "Welcome to the Online Retail System"

	if got != want {
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}
