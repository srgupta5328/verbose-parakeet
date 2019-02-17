package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srgupta5328/verbose-parakeet/helpers"
)

func main() {
	run()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Welcome to Employee Catalog"
	helpers.RespondWithJson(w, 200, msg)
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
