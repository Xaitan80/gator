package main

import (
	"context"
	"fmt"

	"github.com/xaitan80/gator/internal/database" // adjust import path
)

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("must be logged in to run this command")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("must be logged in to run this command: %w", err)
		}

		return handler(s, cmd, user)
	}
}
