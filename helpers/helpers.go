package helpers

import (
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	uuid := uuid.NewV1()
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
