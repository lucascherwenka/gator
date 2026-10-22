package main

import "github.com/lucascherwenka/gator/internal/config"

type state struct {
	cfg *Config 
}

type command struct {
	name  string
	args  []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
