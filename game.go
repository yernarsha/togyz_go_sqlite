package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type TogyzGame struct {
	Id        int
	WhiteName string
	BlackName string
	Result    string
	Event     string
	Date      string
	Site      string
	Notation  string
}

func searchPlayer(db *sql.DB, searchString string) []TogyzGame {

	searchSQL := "SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games WHERE _WhiteName like '%" + searchString + "%' OR _BlackName like '%" + searchString + "%'"
	return searchGames(db, searchSQL)
}

func searchTournament(db *sql.DB, searchString string) []TogyzGame {

	searchSQL := "SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games WHERE _Event like '%" + searchString + "%'"
	return searchGames(db, searchSQL)
}

func searchOpening(db *sql.DB, searchString string) []TogyzGame {

	searchString = "1. " + searchString
	searchSQL := "SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games WHERE _Notation like '" + searchString + "%'"
	return searchGames(db, searchSQL)
}

func searchGames(db *sql.DB, searchSQL string) []TogyzGame {

	rows, err := db.Query(searchSQL)
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	games := make([]TogyzGame, 0)

	for rows.Next() {
		game := TogyzGame{}
		err = rows.Scan(&game.Id, &game.WhiteName, &game.BlackName, &game.Result, &game.Event, &game.Date, &game.Site, &game.Notation)
		if err != nil {
			log.Fatal(err)
		}

		games = append(games, game)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return games
}
