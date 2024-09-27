package main

import "errors"


type cmdInterface interface {
	execute(input string, cfg *config) error
}

type command struct {
	id int
	name string
	description string
}

func (cfg *config) initCommands() error {
	cfg.addNewCommand("create", "creates a new endpoint")
	cfg.addNewCommand("get", "gets all endpoints")
	return nil
}

func (cmd *command) execute(input string, cfg *config) error {
	if cmd.name == "create" {
		commandCreate(input)
	} else if cmd.name == "get" {
		commandGet(input)
	}
	return nil
}

func (cfg *config) getCommands() ([]command, error) {
	return cfg.commands, nil
}

func (cfg *config) getCommand(id int) (command, error) {
	if id > len(cfg.commands) || id < 1 {
		return command{}, errors.New("command not found")
	}

	return cfg.commands[id-1], nil
}

func (cfg *config) addNewCommand(name, description string) {
	id := len(cfg.commands) + 1
	if len(cfg.commands) == 0 {
		id = 1
	}
	cmd := command{
		id: id,
		name: name,
		description: description,
	}
	cfg.commands = append(cfg.commands, cmd)
}

