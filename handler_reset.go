package main

import (
	"fmt"
	"context"
)


func handlerReset(s *state, cmd command) error {
    err := s.db.DeleteUsers(context.Background())
    if err != nil {
        return fmt.Errorf("User delete error: %w", err)
    }
    fmt.Println("Reset Success!")
    return nil
}