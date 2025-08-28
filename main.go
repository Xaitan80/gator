package main

import (
	"fmt"
	"log"

	"github.com/xaitan80/gator/internal/config"
)

func main() {
	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Set current user
	if err := cfg.SetUser("christian"); err != nil {
		log.Fatal(err)
	}

	// Read again to confirm
	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", cfg)
}
