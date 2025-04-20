package main

import (
	"database/sql"
	"fmt"
	"nehnutelnosti-sk/src/internal/parser"
	"nehnutelnosti-sk/src/internal/scrapper"
	"nehnutelnosti-sk/src/internal/store"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	db  *sql.DB
	uri []string
}

func (a *App) CheckUpdated() error {
	// init repo
	repo := store.NewStorage(a.db)
	err := repo.Create()
	if err != nil {
		return err
	}

	for _, p := range a.uri {
		// scrap page
		html, err := scrapper.ScrapWebPage(p)
		if err != nil {
			return err
		}

		// parse html
		parser, err := parser.NewParser(html)
		if err != nil {
			return err
		}

		// get all the flats from the first page
		flats := parser.ParseFlats()
		if len(flats) == 0 {
			return nil
		}

		// check if any title already in DB
		existingFlats, err := repo.SelectExistingFlats(flats)
		if err != nil {
			return err
		}

		fmt.Println(existingFlats)

		// determine what flats are new

		// send notification for any new flats

		// insert new flats
		err = repo.InsertToStore(flats)
		if err != nil {
			return err
		}
	}

	return nil
}
