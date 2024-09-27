package main

import (
	"net/http"
	"os"
)

type config struct {
	httpClient *http.Client
	commands []command
}

func main() {
	cfg := config{
		httpClient: &http.Client{},
		commands: []command{},
	}
	cfg.initCommands()
	repl(os.Stdin, &cfg)
}
