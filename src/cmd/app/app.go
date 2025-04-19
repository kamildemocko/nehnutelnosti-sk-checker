package main

import (
	"fmt"
	"log"
	"nehnutelnosti-sk/src/internal/scrapper"
)

type App struct {
	uri []string
}

func (a *App) CheckUpdated() {
	for _, p := range a.uri {
		text, err := scrapper.ScrapWebPage(p)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(text)
	}
}
