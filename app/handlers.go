package app

import (
	"net/http"

	"github.com/srgupta5328/verbose-parakeet/helpers"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Welcome to Employee Catalog"
	helpers.RespondWithJson(w, 200, msg)
}
