package store

import (
	"context"
	"database/sql"
	"fmt"
	"nehnutelnosti-sk/src/internal/parser"
	"time"
)

type SqliteStorage struct {
	db *sql.DB
}

func (ss *SqliteStorage) Create() error {
	fmt.Println("creating db")

	query := `CREATE TABLE IF NOT EXISTS seen (
	title TEXT PRIMARY KEY,
	created TEXT,
	address TEXT,
	size TEXT,
	area NUMBER,
	price NUMBER,
	link TEXT
	)`

	indexQueries := []string{
		"CREATE INDEX IF NOT EXISTS IDX_title ON seen(title)",
		"CREATE INDEX IF NOT EXISTS IDX_created ON seen(created)",
		"CREATE INDEX IF NOT EXISTS IDX_price ON seen(price)",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ss.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	for _, query := range indexQueries {
		_, err := ss.db.ExecContext(ctx, query)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ss *SqliteStorage) SelectExistingFlats(flats []parser.Flat) ([]*parser.Flat, error) {
	return nil, nil
}

func (ss *SqliteStorage) InsertToStore(flats []parser.Flat) error {
	return nil
}
