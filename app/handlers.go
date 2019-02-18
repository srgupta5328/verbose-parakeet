package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/srgupta5328/verbose-parakeet/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Employee Catalog")
}

func ReadEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{FirstName: "Rohan", LastName: "Gupta", ID: "1", Email: "TestEmail@gmail.dom", RoleName: "Software Engineer 1"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

}

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	e := model.Employee{}
	db, _ := InitDB()
	e.CreateEmployee(db)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
