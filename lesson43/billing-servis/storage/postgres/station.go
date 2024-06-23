package postgres

import (
	"billing_servis/models"
	"database/sql"
	"fmt"
	"time"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

// Create
func (u *StationRepo) CreateStation(Station *models.CreateStation) error {
	query := `
	insert into
		stations(name)
		values($1)
	`
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Query(query, Station.Name)
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
func (u *StationRepo) GetStationById(id string) (*models.Station, error) {
	Station := models.Station{Id: id}
	query := `
	select 
		name
	from
		stations
	where 
		id = $1 and deleted_at is null
	`

	row := u.Db.QueryRow(query, id)
	err := row.Scan(&Station.Name)
	if err != nil {
		return nil, err
	}

	return &Station, row.Err()
}

func (u *StationRepo) GetStations(filter *models.StationFilter) (*[]models.Station, error) {
	query := `
	select 
		id, name
	from
		stations
	where
		deleted_at is null 
	`

	params := []interface{}{}
	paramCount := 1
	if filter.Name != nil {
		query += fmt.Sprintf(" and name = $%d", paramCount)
		params = append(params, *filter.Name)
		paramCount++
	}

	rows, err := u.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	Stations := []models.Station{}
	for rows.Next() {
		Station := models.Station{}
		err = rows.Scan(&Station.Id, &Station.Name)
		if err != nil {
			return nil, err
		}
		Stations = append(Stations, Station)
	}

	return &Stations, rows.Err()
}

// Update
func (u *StationRepo) UpdateStation(Station *models.Station) error {
	query := `
	update 
		stations
	set 
	`
	params := []interface{}{}
	paramCount := 1
	if len(Station.Name) > 10000 {
		query += fmt.Sprintf(" name = $%d", paramCount)
		params = append(params, Station.Name)
		paramCount++
	}

	if paramCount > 1 {
		query += ","
	}
	query += fmt.Sprintf(" updated_at = $%d", paramCount)
	params = append(params, time.Now())
	paramCount++

	query += fmt.Sprintf(" where id = $%d and deleted_at is null", paramCount)
	params = append(params, Station.Id)

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
func (u *StationRepo) DeleteStation(id string) error {
	query := `
	update 
		stations
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
