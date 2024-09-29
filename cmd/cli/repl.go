package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func repl(r io.Reader, cfg *config) {
	scanner := bufio.NewScanner(r)
	fmt.Println(startMessage(cfg))
	for {
		fmt.Print("Mock >> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "help" {
			fmt.Println(startMessage(cfg))
			continue
		}

		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid command")
			continue
		}
		cmd, err := cfg.getCommand(id)
		if err != nil {
			fmt.Println("invalid command")
			continue
		}

		err = cmd.callback(cfg, input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func startMessage(cfg *config) string {
	msg := "Welcome to the API Mocking Service, type help to get a list of available commands\nplease choose an option from the list below:\n"
	options := cfg.getCommands()
	for i := 0; i < len(options); i++ {
		opt := options[i]
		msg += fmt.Sprintf("%d. %s: %s\n", opt.id, opt.name, opt.description)
	}

	return msg
}
