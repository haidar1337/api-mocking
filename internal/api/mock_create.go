package api

import (
	"encoding/json"
	"net/http"

	db "github.com/haidar1337/api-mocking/internal/database"
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
