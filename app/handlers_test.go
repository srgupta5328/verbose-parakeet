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
	want := "Welcome to the Employee Catalog"

	if got != want {
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}

func TestReadEmployeeHandler(t *testing.T) {
	t.Run("Reading the employee from the JSON file", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:8080/test/employee", nil)
		res := httptest.NewRecorder()
		ReadEmployeeHandler(res, req)

		if res.Body.String() == "" {
			t.Errorf("Error reading the JSON test employee")
		}
	})
}
