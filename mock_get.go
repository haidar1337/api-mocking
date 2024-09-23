package main

import (
	"net/http"

	db "github.com/haidar1337/api-mocking/internal/database"
)

func handleGetEndpoints(w http.ResponseWriter, req *http.Request) {
	db, err := db.NewDB("./database.json")
	if err != nil {
		sendErrorResponse(w, 500, "failed to load database")
		return
	}

	endpoints, err := db.GetEndpoints()
	if len(endpoints) < 1 {
		sendErrorResponse(w, 500, "no endpoitns were created yet")
		return
	}
	if err != nil {
		sendErrorResponse(w, 500, "failed to get endpoints")
		return
	}

	sendJSONResponse(w, 200, endpoints)
}

func handleGetEndpoint(w http.ResponseWriter, req *http.Request) {
	id, err := checkIdPathValue(req, "endpointId")
	if err != nil {
		sendErrorResponse(w, 400, err.Error())
		return
	}

	db, err := db.NewDB("./database.json")
	if err != nil {
		sendErrorResponse(w, 500, "failed to load database")
		return
	}

	endpoint, err := db.GetEndpointById(id)
	if err != nil {
		sendErrorResponse(w, 404, err.Error())
		return
	}

	sendJSONResponse(w, 200, endpoint)
}