package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	stuff, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%v", stuff)
	return nil
}

