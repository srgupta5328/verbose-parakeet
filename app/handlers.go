package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/srgupta5328/verbose-parakeet/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Employee Catalog")
}

func ReadEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	e := model.Employee{}
	b, _ := ioutil.ReadFile("../test/employee.json")
	json.Unmarshal(b, &e)
	res, _ := json.Marshal(e)
	w.Write(res)

}
