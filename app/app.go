package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
	"github.com/srgupta5328/verbose-parakeet/helpers"
)

func Initialize() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/listings/latest", LatestListingHandler).Methods("GET")
	return router
}

func Run() {
	router := Initialize()
	fmt.Println("Running Service http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+helpers.Port, router))

}
