package main

import (
	"context"
	"fmt"
	"github.com/cahenrichs/Gator/internal/database"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:		uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:		name,
})
if err != nil {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		fmt.Println("user already exists")
		os.Exit(1)
	}
	return fmt.Errorf("couldn't create user %v", err)
}

err = s.cfg.SetUser(user.Name)
if err != nil {
	return fmt.Errorf("couldn't set the current user %s", err)
}


fmt.Println("User was created")
fmt.Printf("created user: %+v\n", user)
return nil

}

func handlerListUsers(s *State, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not list user %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil

}


