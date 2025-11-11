package main

import (
"github.com/cahenrichs/Gator/internal/config"
"log"
"fmt"
"os"
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

	s := &State{cfg: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*State, command) error),
	}

	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}