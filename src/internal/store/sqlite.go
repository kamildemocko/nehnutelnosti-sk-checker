package store

import (
	"database/sql"
	"nehnutelnosti-sk/src/internal/parser"
)

type SqliteStorage struct {
	db *sql.DB
}

func (ss *SqliteStorage) AreInStore() (bool, error) {
	return false, nil
}

func (ss *SqliteStorage) InsertToStore(flats []parser.Flat) error {
	return nil
}
