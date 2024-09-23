package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}


	InitializeMockAPI(mux)
	

	log.Printf("Server started on port: %s", port)
	log.Fatal(server.ListenAndServe())
}