package main

import (
	"context"
	"fmt"
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
