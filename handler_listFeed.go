package main

import (
	"fmt"
	"context"
)


func handlerListFeeds(s *state, cmd command) error {
    feeds, err := s.db.GetFeeds(context.Background())
    if err != nil {
        return fmt.Errorf("Feed list error: %w", err)
    }

    for _, feed := range feeds {

		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("User get error: %w", err)
		}

		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}
    return nil
}