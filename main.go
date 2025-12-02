package main

import (
	"context"
"database/sql"
"github.com/cahenrichs/Gator/internal/config"
"github.com/cahenrichs/Gator/internal/database"
"log"
"fmt"
"os"
_ "github.com/lib/pq"
)

type State struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	defer db.Close()
	dbQueries := database.New(db)	

	programState := &State{
		db: dbQueries,
		cfg: &cfg,
	}


	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	// fmt.Printf("%+v\n", cfg)

	// s = &State{cfg: &cfg}

	cmds := commands{
		registeredCommands: make(map[string]func(*State, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListFeedFollows))
	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

	// if err := cmds.run(s, cmd); err != nil {
		// fmt.Println(err)
		// os.Exit(1)
	// }

}
func middlewareLoggedIn(handler func(s *State, cmd command, user database.User) error) func(*State, command) error {
	return func(s *State, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
