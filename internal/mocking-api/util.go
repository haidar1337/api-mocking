package mockingapi

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	type ErrorStruct struct {
		Error string `json:"error"`
	}

	sendJSONResponse(w, code, ErrorStruct{
		Error: message,
	})
}

func sendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("content-type", "application/json")
	dat, err := json.Marshal(&payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(dat)
}