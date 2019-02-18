package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	res := httptest.NewRecorder()

	HomeHandler(res, req)

	got := res.Body.String()
	want := "Welcome to the Coin Market Application in Go"

	if got != want {
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}
