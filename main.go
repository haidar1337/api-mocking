package main

import (
	"log"
	"net/http"

	mockingapi "github.com/haidar1337/api-mocking/internal/mocking-api"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}


	mockingapi.InitializeMockAPI(mux)
	

	log.Printf("Server started on port: %s", port)
	log.Fatal(server.ListenAndServe())

}