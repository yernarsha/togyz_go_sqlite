package main

import (
	"database/sql"
	"fmt"
	"log"
    "os"
	"github.com/dixonwille/wmenu/v5"
)

func main() {

    fmt.Println("Welcome to the Modern TogyzKumalak!\n")
	db, err := sql.Open("sqlite3", "./y.sqlite")
	checkErr(err)
	defer db.Close()

	menu := wmenu.NewMenu("Enter selection:")
	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Show all games", 0, true, nil)
	menu.Option("Filter by Player", 1, false, nil)
	menu.Option("Filter by Tournament", 2, false, nil)
	menu.Option("Filter by Opening", 3, false, nil)
	menu.Option("Exit", 4, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch {

	case opts[0].Value == 0 || opts[0].Value == 1 || opts[0].Value == 2 || opts[0].Value == 3:
        var input string
        var games []TogyzGame
        
        if opts[0].Value == 0 {
            searchSQL := "SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games"
    		games = searchGames(db, searchSQL)
            
        } else if opts[0].Value == 1 {
            fmt.Print("Enter a player name: ")
            fmt.Scan(&input)
    		games = searchPlayer(db, input)
            
        } else if opts[0].Value == 2 {
    		fmt.Print("Enter a tournament name: ")
            fmt.Scan(&input)
    		games = searchTournament(db, input)
            
        } else if opts[0].Value == 3 {
    		fmt.Print("Enter an opening (1-9): ")
            fmt.Scan(&input)
    		games = searchOpening(db, input)            
        }

		for _, game := range games {
			fmt.Printf("\n%d. %s - %s %s, %s, %s, %s",
                       game.Id, game.WhiteName, game.BlackName, game.Result, game.Event, game.Date, game.Site)
		}
		fmt.Printf("\n\nFound %v games\n", len(games))

	case opts[0].Value == 4:
		fmt.Println("Come back again!")
		os.Exit(0)
    default:
        os.Exit(0)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
