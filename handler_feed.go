package main

import (
	"context"
	"fmt"

	"github.com/xaitan80/gator/internal/database"
)

// handlerAgg fetches a feed and prints it
func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	// Print the entire RSSFeed struct
	fmt.Printf("%+v\n", feed)
	return nil
}

// addFeed
func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Printf("Feed created:\n * ID: %v\n * Name: %v\n * URL: %v\n", feed.ID, feed.Name, feed.Url)
	return nil
}
func handlerFeeds(s *state, cmd command) error {
	// Ensure no arguments
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s takes no arguments", cmd.Name)
	}

	// Fetch all feeds from DB
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch feeds: %w", err)
	}

	// Print feed name and the user who created it, each on a separate line
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.UserName)
	}

	return nil
}
