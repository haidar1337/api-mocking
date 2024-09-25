package api

import (
	"net/http"

	db "github.com/haidar1337/api-mocking/internal/database"
)

func handleMockDelete(w http.ResponseWriter, req *http.Request) {
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

	err = db.DeleteEndpoint(id)
	if err != nil {
		sendErrorResponse(w, 404, err.Error())
		return
	}

	sendJSONResponse(w, 204, struct{}{})
}
