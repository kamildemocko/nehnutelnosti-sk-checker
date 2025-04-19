package main

import (
	"fmt"
	"log"
	"nehnutelnosti-sk/src/internal/parser"
	"nehnutelnosti-sk/src/internal/scrapper"
)

type App struct {
	uri []string
}

func (a *App) CheckUpdated() {
	for _, p := range a.uri {
		html, err := scrapper.ScrapWebPage(p)
		if err != nil {
			// todo
			log.Println(err)
			continue
		}

		parser, err := parser.NewParser(html)
		if err != nil {
			// todo
			continue
		}

		// todo
		flats := parser.ParseFlats()
		for _, flat := range flats {
			fmt.Println(flat)
		}
	}
}
