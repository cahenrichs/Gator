package main 

import (
"context"
"fmt"
"github.com/cahenrichs/Gator/internal/database"
"time"
"github.com/google/uuid"
)

func handlerFollow(s *State, cmd command, user database.User) error {
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

	feedFollowRow, err := s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:	feed.ID,
	})

	fmt.Println("Feed follow created")
	fmt.Println(feedFollowRow.UserName, feedFollowRow.FeedName)
	return nil
}

func handlerListFeedFollows(s *State, cmd command, user database.User) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get the feed follows %s", err)
	}

	if len(feedFollow) == 0 {
		fmt.Println("No feed follows found for this user")
		return nil
	}

	fmt.Printf("feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollow {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil

}
func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:			%s\n", username)
	fmt.Printf("* Feed:			%s\n", feedname)
}
