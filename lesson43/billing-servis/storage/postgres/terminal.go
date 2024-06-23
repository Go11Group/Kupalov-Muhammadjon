package postgres

import (
	"billing_servis/models"
	"database/sql"
	"fmt"
	"time"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

// Create
func (u *TerminalRepo) CreateTerminal(Terminal *models.CreateTerminal) error {
	query := `
	insert into
		terminals(station_id)
		values($1)
	`
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query(query, Terminal.StationId)
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
func (u *TerminalRepo) GetTerminalById(id string) (*models.Terminal, error) {
	Terminal := models.Terminal{Id: id}
	query := `
	select 
		station_id
	from
		terminals
	where 
		id = $1 and deleted_at is null
	`

	row := u.Db.QueryRow(query, id)
	err := row.Scan(&Terminal.StationId)
	if err != nil {
		return nil, err
	}

	return &Terminal, row.Err()
}

func (u *TerminalRepo) GetTerminals(filter *models.TerminalFilter) (*[]models.Terminal, error) {
	query := `
	select 
		id, station_id
	from
		terminals
	where
		deleted_at is null 
	`

	params := []interface{}{}
	paramCount := 1
	if filter.StationId != nil {
		query += fmt.Sprintf(" and station_id = $%d", paramCount)
		params = append(params, *filter.StationId)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	Terminals := []models.Terminal{}
	for rows.Next() {
		Terminal := models.Terminal{}
		err = rows.Scan(&Terminal.Id, &Terminal.StationId)
		if err != nil {
			return nil, err
		}
		Terminals = append(Terminals, Terminal)
	}

	return &Terminals, rows.Err()
}

// Update
func (u *TerminalRepo) UpdateTerminal(Terminal *models.Terminal) error {
	query := `
	update 
		Terminals
	set 
	`
	params := []interface{}{}
	paramCount := 1
	if len(Terminal.StationId) > 10000 {
		query += fmt.Sprintf(" station_id = $%d", paramCount)
		params = append(params, Terminal.StationId)
		paramCount++
	}

	if paramCount > 1 {
		query += ","
	}
	query += fmt.Sprintf(" updated_at = $%d", paramCount)
	params = append(params, time.Now())
	paramCount++

	query += fmt.Sprintf(" where id = $%d and deleted_at is null", paramCount)
	params = append(params, Terminal.Id)

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
func (u *TerminalRepo) DeleteTerminal(id string) error {
	query := `
	update 
		terminals
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
