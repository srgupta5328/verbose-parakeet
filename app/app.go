package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/srgupta5328/verbose-parakeet/helpers"
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
	log.Fatal(http.ListenAndServe(":"+helpers.Port, router))

	return nil

}

func InitDB() (*scribble.Driver, error) {
	db, err := scribble.New("./employees", nil)
	if err != nil {
		fmt.Println("Error initializing the scribble db")
	}

	return db, nil
}
