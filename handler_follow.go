package main

import (
	"context"
	"fmt"

	"github.com/xaitan80/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feedURL := cmd.Args[0]
	currentUser := s.cfg.CurrentUserName

	// 1. Get current user from DB
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("couldn't find current user: %w", err)
	}

	// 2. Get feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	// 3. Create feed follow
	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	// 4. Print confirmation
	fmt.Printf("%s now follows %s\n", follow.UserName, follow.FeedName)
	return nil
}
