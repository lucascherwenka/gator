package main

import (
	"fmt"
	"log"

	"github.com/lucascherwenka/gator/internal/config"
)

func main() {
	cfg, err := state{config.Read()}
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
}