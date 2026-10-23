package main

import (
	"github.com/lucascherwenka/gator/internal/config"
	"errors"
	"github.com/lucascherwenka/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries 
}

type command struct {
	name  string
	args  []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandMap[cmd.name]
	if !ok {
		return errors.New("Not a known command")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
