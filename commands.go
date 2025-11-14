package main

import (
"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {

	registeredCommands map[string]func(*State, command) error
}

func (c *commands) run(s *State, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*State, command) error) {
	c.registeredCommands[name] = f
}
