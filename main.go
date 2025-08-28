package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xaitan80/gator/internal/cli"
	"github.com/xaitan80/gator/internal/config"
)

func main() {
	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize state
	appState := &cli.State{
		Config: &cfg,
	}

	// Initialize commands and register handlers
	cmds := &cli.Commands{}
	cmds.Register("login", cli.HandlerLogin)

	// Check command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments")
		os.Exit(1)
	}

	// Split command name and arguments
	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// Run the command
	if err := cmds.Run(appState, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
