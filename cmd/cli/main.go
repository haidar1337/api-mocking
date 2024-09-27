package main

import (
	"net/http"
	"os"
)

type config struct {
	baseURL    string
	httpClient *http.Client
	commands   []command
}

func main() {
	cfg := config{
		httpClient: &http.Client{},
		commands:   []command{},
		baseURL:    baseURL,
	}
	cfg.initCommands()
	repl(os.Stdin, &cfg)
}
