package main

import "errors"

type command struct {
	id          int
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func (cfg *config) initCommands() error {
	cfg.addNewCommand("create", "creates a new endpoint", commandCreate)
	cfg.addNewCommand("get", "gets all endpoints", commandGet)
	return nil
}

func (cfg *config) getCommands() []command {
	return cfg.commands
}

func (cfg *config) getCommand(id int) (command, error) {
	if id > len(cfg.commands) || id < 1 {
		return command{}, errors.New("command not found")
	}

	return cfg.commands[id-1], nil
}

func (cfg *config) addNewCommand(name, description string, callback func(cfg *config, args ...string) error) {
	id := len(cfg.commands) + 1
	if len(cfg.commands) == 0 {
		id = 1
	}
	cmd := command{
		id:          id,
		name:        name,
		description: description,
		callback:    callback,
	}
	cfg.commands = append(cfg.commands, cmd)
}
