package cli

import (
	"errors"
	"fmt"

	"github.com/xaitan80/gator/internal/config"
)

// exported State struct
type State struct {
	Config *config.Config
}

// exported Command struct
type Command struct {
	Name string
	Args []string
}

// exported Commands struct
type Commands struct {
	Handlers map[string]func(*State, Command) error
}

// exported Register method
func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*State, Command) error)
	}
	c.Handlers[name] = f
}

// exported Run method
func (c *Commands) Run(s *State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}

// exported login handler
func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("username argument is required")
	}

	username := cmd.Args[0]

	if err := s.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("Current user set to: %s\n", username)
	return nil
}
