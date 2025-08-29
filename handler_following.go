package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	// Ensure no extra arguments
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s takes no arguments", cmd.Name)
	}

	// Get current username from config
	currentUser := s.cfg.CurrentUserName
	if currentUser == "" {
		return fmt.Errorf("no current user logged in")
	}

	// Get user record to fetch ID
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't find current user: %w", err)
	}

	// Fetch feed follows for this user
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to fetch feed follows: %w", err)
	}

	// Print the names of the feeds
	for _, ff := range follows {
		fmt.Println(ff.FeedName)
	}

	return nil
}
