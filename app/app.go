package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/srgupta5328/verbose-parakeet/helpers"
)

type App struct {
	DB     *sql.DB
	router *mux.Router
}

func Initialize() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/test/employee", ReadEmployeeHandler).Methods("GET")
	return router
}

func InitDB(username, password, dbname string) *sql.DB {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", username, password, dbname)

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (a *App) Run() error {
	a.router = Initialize()
	a.DB = InitDB(helpers.Username, helpers.Password, helpers.DbName)
	fmt.Println("Running Service http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+helpers.Port, a.router))

	return nil

}
