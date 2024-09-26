package api

import "net/http"

func InitializeMockAPI(mux *http.ServeMux) {

	mux.HandleFunc("GET /mock/endpoints", handleGetEndpoints)
	mux.HandleFunc("GET /mock/endpoints/{endpointId}", handleGetEndpoint)
	mux.HandleFunc("POST /mock/endpoints", handleMockCreate)
	mux.HandleFunc("PUT /mock/endpoints/{endpointId}", handleMockUpdate)
	mux.HandleFunc("DELETE /mock/endpoints/{endpointId}", handleMockDelete)

	mux.HandleFunc("GET /mock/{endpointId}", handleMocking)
}
