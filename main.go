package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/gorilla/mux"
)

func main() {
	run()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Employee Catalog")
}

func Initialize() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	return router
}

func run() error {
	router := Initialize()
	fmt.Println("Running Service http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	return nil

}

func generateUUID() string {
	uuid, _ := uuid.NewV1()
	id := uuid.String()
	return id
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type Error struct {
	msg  string
	code int
}

func ResponseWithError(w http.ResponseWriter, e Error) {
	RespondWithJson(w, e.code, e.msg)
}

type Employee struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	RoleName  string `json:"role_name,omitempty"`
}

func (e *Employee) createEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) updateEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) deleteEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}

func (e *Employee) getEmployee(db *sql.DB) error {
	return errors.New("Not implemented yet")
}
