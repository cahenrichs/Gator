package main 

import (
"context"
"fmt"
"github.com/cahenrichs/Gator/internal/database"
"time"
)

func handlerFollow (s *State, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get the feed %w", err)
	}

	feedFollowRow, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})

	fmt.Println("Feed follow created")
	fmt.Println(feedFollowRow.UserName, feedFollowRow.FeedName)
	return nil
}