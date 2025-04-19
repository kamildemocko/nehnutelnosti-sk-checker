package store

import (
	"database/sql"
	"fmt"
	"nehnutelnosti-sk/src/internal/parser"
)

type SqliteStorage struct {
	db *sql.DB
}

func (ss *SqliteStorage) Create() error {
	fmt.Println("creating db")
	return nil
}

func (ss *SqliteStorage) SelectExistingFlats(flats []parser.Flat) ([]*parser.Flat, error) {
	return nil, nil
}

func (ss *SqliteStorage) InsertToStore(flats []parser.Flat) error {
	return nil
}
