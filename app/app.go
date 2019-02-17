package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/test/employee", ReadEmployeeHandler).Methods("GET")
	return router
}

func Run() error {
	router := Initialize()
	fmt.Println("Running Service http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	return nil

}
