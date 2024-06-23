package models

type Transaction struct {
	Id              string `json:"id"`
	CardId          string `json:"card_id"`
	Amount          int    `json:"amount"`
	TerminalId      string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
	Time
}

type CreateTransaction struct {
	CardId          string `json:"card_id"`
	Amount          int    `json:"amount"`
	TerminalId      string `json:"terminal_id"`
	TransactionType string `json:"transaction_type"`
}

type TransactionFilter struct {
	CardId          *string `json:"card_id"`
	Amount          *int    `json:"amount"`
	TerminalId      *string `json:"terminal_id"`
	TransactionType *string `json:"transaction_type"`
}
