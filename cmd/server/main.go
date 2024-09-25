package main

import (
	"log"
	"net/http"

	"github.com/haidar1337/api-mocking/internal/api"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	api.InitializeMockAPI(mux)

	log.Printf("Server started on port: %s", port)
	log.Fatal(server.ListenAndServe())
}
