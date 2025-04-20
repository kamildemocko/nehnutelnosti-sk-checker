package store

import (
	"database/sql"
	"nehnutelnosti-sk/src/internal/parser"
)

type Storage interface {
	Create() error
	SelectExistingFlats(flats []*parser.Flat) ([]*parser.Flat, error)
	InsertToStore(flats []*parser.Flat) error
}

func NewStorage(db *sql.DB) Storage {
	return &SqliteStorage{
		db: db,
	}
}
