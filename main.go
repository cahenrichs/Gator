package main

import (
"github.com/cahenrichs/Gator/internal/config"
"log"
"fmt"
)

type State struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cfg.SetUser("Clint")
	if err != nil {
		return 
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
}