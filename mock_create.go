package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	db "github.com/haidar1337/api-mocking/internal/database"
)

type RequestMethod string

const (
	RequestMethodGET  	 	RequestMethod = "GET"
	RequestMethodPOST 		RequestMethod = "POST"
	RequestMethodDELETE  RequestMethod = "DELETE"
	RequestMethodPUT   	 RequestMethod = "PUT"
	RequestMethodPATCH    RequestMethod = "PATCH"
)

func handleMockCreate(w http.ResponseWriter, req *http.Request) {
	endpoint := db.MockEndpoint{}
	err := json.NewDecoder(req.Body).Decode(&endpoint)
	if err != nil {
		sendErrorResponse(w, 400, "bad request; provide a correct request body")
		return
	}
	if endpoint.Endpoint == "" {
		sendErrorResponse(w, 400, "endpoint cannot be empty")
		return
	}
	if endpoint.Response == (db.MockEndpointResponse{}) {
		sendErrorResponse(w, 400, "bad request; provide a response object")
		return
	}
	_, err = sanitizeMethodInput(endpoint.Method)
	if err != nil {
		sendErrorResponse(w, 400, err.Error())
		return
	}
	

	db, err := db.NewDB("./database.json")
	if err != nil {
		sendErrorResponse(w, 500, "failed to load database")
		return
	}

	ep, err := db.CreateMockEndpoint(endpoint)
	if err != nil {
		sendErrorResponse(w, 500, "failed to create endpoint")
		return
	}

	sendJSONResponse(w, 201, ep)
}

func sanitizeMethodInput(method string) (string, error) {
	cleaned := strings.TrimSpace(strings.ToUpper(method))
	fmt.Println(cleaned, string(RequestMethodDELETE))
	if cleaned == string(RequestMethodPOST) || cleaned == string(RequestMethodGET) || cleaned == string(RequestMethodPUT) || cleaned == string(RequestMethodPATCH) || cleaned == string(RequestMethodDELETE) {
		return cleaned, nil
	}

	return "", errors.New("invalid method; method has to be one of: POST, GET, DELETE, PUT, PATCH")
}