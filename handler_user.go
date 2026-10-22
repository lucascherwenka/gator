package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires a name")
	}
	s.cfg.CurrentUserName = cmd.args[0]
	fmt.Printf("User has been set to %v", cmd)
	return nil
}