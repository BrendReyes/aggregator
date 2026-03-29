package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/brendreyes/aggregator/internal/database"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feed_name := cmd.Args[0]
	url := cmd.Args[1]

    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("User does not exist: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feed_name,
		Url:	   url, 
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("feed creation error: %w", err)
	}

	fmt.Println(feed)

    
    return nil
}
