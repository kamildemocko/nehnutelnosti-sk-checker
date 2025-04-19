package store

import (
	"database/sql"
	"nehnutelnosti-sk/src/internal/parser"
)

type Storage interface {
	AreInStore() (bool, error)
	InsertToStore(flats []parser.Flat) error
}

func NewStorage(db *sql.DB) Storage {
	return &SqliteStorage{
		db: db,
	}
}
