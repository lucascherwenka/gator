package main

import (
	"errors"
	"fmt"
	"time"
	"context"
	"github.com/lucascherwenka/gator/internal/database"
	"github.com/google/uuid"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Login requires a name")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("User not found")
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User has been set to %v\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Registration failed!")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		cmd.args[0],
	})
	if err != nil {
		return errors.New("CreateUser failed")
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Printf("User has been created and set to %v\n", cmd.args[0])

	log.Printf("%+v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
		fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}