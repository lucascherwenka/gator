package main

import (
	"log"
	"os"
	"github.com/lucascherwenka/gator/internal/config"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/lucascherwenka/gator/internal/database"
	"context"
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
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(addFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnFollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))


	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments")
	}

	input := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(s, input)
	if err != nil {
		log.Fatalf("%v", err)
	}	
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}