package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)



func repl(r io.Reader, cfg *config) {
	scanner := bufio.NewScanner(r)
	fmt.Println(startMessage(cfg))
	for {
		fmt.Print("Mock >> ")
		scanner.Scan()
		input := scanner.Text()

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

		cmd.execute(input, cfg)
	}
}

func startMessage(cfg *config) string {
	welcome := "Welcome to the API Mocking Service, please choose an option from the list below:\n"
	options, err := cfg.getCommands()
	if err != nil {
		return welcome
	}
	msg := fmt.Sprintf(welcome)
	for i := 0; i < len(options); i++ {
		opt := options[i]
		msg += fmt.Sprintf("%d. %s. %s\n", opt.id, opt.name, opt.description)
	}

	return msg
}
