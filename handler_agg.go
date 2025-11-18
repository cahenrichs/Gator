package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *State, cmd command) error {
	ctx := context.Background()
	res, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}