package api

import (
	"net/http"
	"time"

	db "github.com/haidar1337/api-mocking/internal/database"
)

func handleMocking(w http.ResponseWriter, req *http.Request) {
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

	simulateDelay(endpoint.Delay)
	sendJSONResponse(w, endpoint.Response.StatusCode, endpoint)
}

func simulateDelay(delay uint) {
	if delay != 0 {
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}
