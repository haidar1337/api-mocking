package main

import (
	"log"
	"net/http"
)

type config struct {
	mux *http.ServeMux
}

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	cfg := config{
		mux: mux,
	}
	cfg.InitializeMockAPI(mux)

	log.Printf("Server started on port: %s", port)
	log.Fatal(server.ListenAndServe())
}
