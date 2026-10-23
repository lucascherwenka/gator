package main

import (
	"log"
	"os"
	"github.com/lucascherwenka/gator/internal/config"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/lucascherwenka/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("%v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	dbQueries := database.New(db)

	s := &state{cfg: &cfg, db: dbQueries}
	
	cmds := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerUsers)
	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments")
	}

	input := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(s, input)
	if err != nil {
		log.Fatalf("%v", err)
	}	
}