package handler

import (
	"billing_servis/storage/postgres"
	"database/sql"
)

type handler struct{
	CardRepo *postgres.CardRepo
	StationRepo *postgres.StationRepo
	TerminalRepo *postgres.TerminalRepo
	TransactionRepo *postgres.TransactionRepo
}

func newHandler(db *sql.DB) *handler{
	c := postgres.NewCardRepo(db)
	s := postgres.NewStationRepo(db)
	t := postgres.NewTerminalRepo(db)
	tr := postgres.NewTransactionRepo(db)

	return &handler{
		CardRepo: c,
		StationRepo: s,
		TerminalRepo: t,
		TransactionRepo: tr,
	}
}