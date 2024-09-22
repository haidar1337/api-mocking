package mockingapi

import "net/http"

func InitializeMockAPI(mux *http.ServeMux) {
	
	mux.HandleFunc("GET /mock/endpoints", handleMockGet)
	mux.HandleFunc("POST /mock/endpoints", handleMockCreate)
	mux.HandleFunc("PUT /mock/endpoints/{endpoint}", handleMockUpdate)
	mux.HandleFunc("DELETE /mock/endpoints/{endpoint}", handleMockDelete)

	mux.HandleFunc("/mock/{endpoint}", handleMocking)
}