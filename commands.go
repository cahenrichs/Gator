package main

import (
"fmt"
"errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {

	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

}

func (c *commands) register(name string, f func(*state, command) error) {
	
}

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) < 1 {
        return fmt.Errorf("usage: %s <username>", cmd.name)

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

func handlerLogin(s *State, cmd command) error {
    if len(cmd.Args) < 1 {
        return fmt.Errorf("usage: %s <username>", cmd.Name)

    }
    if s.cfg == nil {
        return fmt.Errorf("state config not initialized")
    }

    username := cmd.args[0]

    username := cmd.Args[0]

    if err := s.cfg.SetUser(username); err != nil {
        return fmt.Errorf("could not set current user: %w", err)
    }

    fmt.Printf("User set to '%s'\n", username)
    return nil
}
/*
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	username := cmd.Args[0]
	if s.Config == nil {
		return fmt.Errorf("state config not initialized")
	}

	err := s.cfg.SetUser(user.username)
	if err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	s.Config.User = username
	fmt.Printf("User set to '%s'\n", username)
	return nil
}*/

