package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/brendreyes/aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User does not exist: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %v <name>", cmd.Name)
    }

    name := cmd.Args[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

    user, err := s.db.CreateUser(context.Background(), params)
    if err != nil {
		return fmt.Errorf("user creation error: %w", err)
	}

    err = s.cfg.SetUser(name)
    if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println(user)
    return nil
}


