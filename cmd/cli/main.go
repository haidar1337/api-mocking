package main

import (
	"net/http"
	"os"
)

type config struct {
	httpClient *http.Client
	commands   []command
	baseURL    string
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
