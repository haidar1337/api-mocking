package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

func checkIdPathValue(req *http.Request, path string) (int, error) {
	in := req.PathValue("endpointId")
	if in == "" {
		return 0, errors.New("please provide an id")
	}

	id, err := strconv.Atoi(in)
	if err != nil {
		return 0, errors.New("please provide a valid id")
	}

	return id, nil
}