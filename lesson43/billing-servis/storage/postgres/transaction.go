package postgres

import (
	"billing_servis/models"
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

// Create
func (u *TransactionRepo) CreateTransaction(Transaction *models.CreateTransaction) error {
	if Transaction.TransactionType == "credit" {
		balanceQuery := `
		select 
			sum(case when transaction_type = 'debit' then amount end) - 
			sum(case when transaction_type = 'credit' then amount end) as balance
		from
			transactions
		where 
			card_id = $1
		`
		balance := 0
		row := u.Db.QueryRow(balanceQuery, Transaction.CardId)
		err := row.Scan(&balance)
		if err != nil {
			return err
		}
		if balance-Transaction.Amount < 0 {
			return fmt.Errorf("not enough money to withdraw")
		}
	}

	query := `
	insert into
		transactions(car_id, amount, terminal_id, transaction_type)
		values($1, $2, $3, $4)
	`
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query(query, Transaction.CardId, Transaction.Amount,
		Transaction.TerminalId, Transaction.TransactionType)

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Read
func (u *TransactionRepo) GetTransactionById(id string) (*models.Transaction, error) {
	Transaction := models.Transaction{Id: id}
	query := `
	select 
		car_id, amount, terminal_id, transaction_type
	from
		transactions
	where 
		id = $1 and deleted_at is null
	`

	row := u.Db.QueryRow(query, id)
	err := row.Scan(&Transaction.CardId, &Transaction.Amount, &Transaction.TerminalId, &Transaction.TransactionType)
	if err != nil {
		return nil, err
	}

	return &Transaction, row.Err()
}

func (u *TransactionRepo) GetTransactions(filter *models.TransactionFilter) (*[]models.Transaction, error) {
	query := `
	select 
		id, car_id, amount, terminal_id, transaction_type
	from
		transactions
	where
		deleted_at is null 
	`

	params := []interface{}{}
	paramCount := 1
	if filter.CardId != nil {
		query += fmt.Sprintf(" and car_id = $%d", paramCount)
		params = append(params, *filter.CardId)
		paramCount++
	}
	if filter.Amount != nil {
		query += fmt.Sprintf(" and amount = $%d", paramCount)
		params = append(params, *filter.Amount)
		paramCount++
	}
	if filter.TerminalId != nil {
		query += fmt.Sprintf(" and terminal_id = $%d", paramCount)
		params = append(params, *filter.TerminalId)
		paramCount++
	}
	if filter.TransactionType != nil {
		query += fmt.Sprintf(" and transaction_type = $%d", paramCount)
		params = append(params, *filter.TransactionType)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	Transactions := []models.Transaction{}
	for rows.Next() {
		Transaction := models.Transaction{}
		err = rows.Scan(&Transaction.Id, &Transaction.CardId, &Transaction.Amount, &Transaction.TerminalId)
		if err != nil {
			return nil, err
		}
		Transactions = append(Transactions, Transaction)
	}

	return &Transactions, rows.Err()
}

// Update
func (u *TransactionRepo) UpdateTransaction(Transaction *models.Transaction) error {
	query := `
	update 
		transactions
	set 
	`
	params := []interface{}{}
	paramCount := 1
	if len(Transaction.CardId) > 0 {
		query += fmt.Sprintf(" card_id = $%d", paramCount)
		params = append(params, Transaction.CardId)
		paramCount++
	}
	if Transaction.Amount > 0 {
		if paramCount > 1 {
			query += ","
		}
		query += fmt.Sprintf(" amount = $%d", paramCount)
		params = append(params, Transaction.Amount)
		paramCount++
	}
	if len(Transaction.TerminalId) > 0 {
		if paramCount > 1 {
			query += ","
		}
		query += fmt.Sprintf(" terminal_id = $%d", paramCount)
		params = append(params, Transaction.TerminalId)
		paramCount++
	}

	if paramCount > 1 {
		query += ","
	}
	query += fmt.Sprintf(" updated_at = $%d", paramCount)
	params = append(params, time.Now())
	paramCount++

	query += fmt.Sprintf(" where id = $%d and deleted_at is null", paramCount)
	params = append(params, Transaction.Id)

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(query, params...)
	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affectedRows == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing updated")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// Delete
func (u *TransactionRepo) DeleteTransaction(id string) error {
	query := `
	update 
		transactions
	set
		deleted_at = $1
	where
		id = $2 and deleted_at is null
	`

	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(query, time.Now(), id)
	if err != nil {
		tx.Rollback()
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if affectedRows == 0 {
		tx.Rollback()
		return fmt.Errorf("nothing deleted")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func checkBalance(cardId int, amount int) {

}
