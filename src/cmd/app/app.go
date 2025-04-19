package main

import (
	"fmt"
	"nehnutelnosti-sk/src/internal/parser"
	"nehnutelnosti-sk/src/internal/scrapper"
)

type App struct {
	uri []string
}

func (a *App) CheckUpdated() error {
	for _, p := range a.uri {
		html, err := scrapper.ScrapWebPage(p)
		if err != nil {
			return err
		}

		parser, err := parser.NewParser(html)
		if err != nil {
			return err
		}

		flats := parser.ParseFlats()
		// remove printing
		for _, flat := range flats {
			fmt.Println(flat)
		}

		// check if any title already in DB

		// send notification for any new flats

		// insert any new flats
	}

	return nil
}
