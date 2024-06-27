package postgres

import "database/sql"

type TransportRepo struct {
	DB *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{db}
}

func (t *TransportRepo) Create() {

}
