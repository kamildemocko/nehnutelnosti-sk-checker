package main

import (
	"database/sql"
	"nehnutelnosti-sk/src/internal/uri"
)

func main() {
	uri1, err := uri.NewUrlBuilder().
		WithPlace("kosice").
		WithSize("3-izbove-byty").
		WithArea(50, 100).
		WithPrice(100000, 200000).
		Build()
	if err != nil {
		panic("wrong parameters for uri")
	}

	db, err := initDB()
	if err != nil {
		panic(err)
	}

	app := App{
		uri: []string{
			uri1,
		},
		db: db,
	}

	err = app.CheckUpdated()
	if err != nil {
		panic(err)
	}
}

func initDB() (*sql.DB, error) {
	sqlite, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}

	return sqlite, nil
}
