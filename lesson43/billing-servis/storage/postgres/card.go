package postgres

import (
	"billing_servis/models"
	"database/sql"
	"fmt"
	"time"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

// Create
func (u *CardRepo) CreateCard(Card *models.CreateCard) error {
	query := `
	insert into
		cards(number, user_id)
		values($1, $2)
	`
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query(query, Card.Number, Card.UserId)
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
func (u *CardRepo) GetCardById(id string) (*models.Card, error) {
	Card := models.Card{Id: id}
	query := `
	select 
		number, user_id
	from
		cards
	where 
		id = $1 and deleted_at is null
	`

	row := u.Db.QueryRow(query, id)
	err := row.Scan(&Card.Number, &Card.UserId)
	if err != nil {
		return nil, err
	}

	return &Card, row.Err()
}

func (u *CardRepo) GetCards(filter *models.CardFilter) (*[]models.Card, error) {
	query := `
	select 
		id, number, user_id
	from
		cards
	where
		deleted_at is null 
	`

	params := []interface{}{}
	paramCount := 1
	if filter.Number != nil {
		query += fmt.Sprintf(" and number = $%d", paramCount)
		params = append(params, *filter.Number)
		paramCount++
	}
	if filter.UserId != nil {
		query += fmt.Sprintf(" and user_id = $%d", paramCount)
		params = append(params, *filter.UserId)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	Cards := []models.Card{}
	for rows.Next() {
		Card := models.Card{}
		err = rows.Scan(&Card.Id, &Card.Number, &Card.UserId)
		if err != nil {
			return nil, err
		}
		Cards = append(Cards, Card)
	}

	return &Cards, rows.Err()
}

// Update
func (u *CardRepo) UpdateCard(Card *models.Card) error {
	query := `
	update 
		cards
	set 
	`
	params := []interface{}{}
	paramCount := 1
	if Card.Number > 10000 {
		query += fmt.Sprintf(" number = $%d", paramCount)
		params = append(params, Card.Number)
		paramCount++
	}
	if len(Card.UserId) > 0 {
		if paramCount > 1 {
			query += ","
		}
		query += fmt.Sprintf(" user_id = $%d", paramCount)
		params = append(params, Card.UserId)
		paramCount++
	}

	if paramCount > 1 {
		query += ","
	}
	query += fmt.Sprintf(" updated_at = $%d", paramCount)
	params = append(params, time.Now())
	paramCount++

	query += fmt.Sprintf(" where id = $%d and deleted_at is null", paramCount)
	params = append(params, Card.Id)

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
func (u *CardRepo) DeleteCard(id string) error {
	query := `
	update 
		cards
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
