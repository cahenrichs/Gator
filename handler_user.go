package main

import (
	"context"
	"fmt"
	"github.com/cahenrichs/Gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil

}

func handlerRegister(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParms{
		ID:		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:		name,
})
if err != nil {
	return fmt.Errorf("couldn't create user %v", err)
}

err = s.cfg.SetUser(user.Name)
if err != nil {
	return fmt.Errorf("couldn't set the current user %s", err)
}

fmt.Println("User was created")

return nil

}
