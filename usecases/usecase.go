package usecases

import "database/sql"

type Usecases struct {
	db *sql.DB
}

func NewUsecases(db *sql.DB) *Usecases {
	return &Usecases{db: db}
}