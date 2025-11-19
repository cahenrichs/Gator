package main

import (
	"context"
	"fmt"
	"github.com/cahenrichs/Gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *State, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	
	name := cmd.Args[0]
	url := cmd.Args[1]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		name,
		Url:		url,
		UserID:		user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Feed was created")
	fmt.Printf("created feed: %+v\n", feed)
	return nil
}