package datastore

import (
	"github.com/jmoiron/sqlx"
)

type Storage interface {
}

type DisputeStore struct {
	DB *sqlx.DB
}

func NewDisputeStore(db *sqlx.DB) *DisputeStore {
	return &DisputeStore{
		DB: db,
	}
}
