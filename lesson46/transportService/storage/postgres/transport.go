package postgres

import (
	"database/sql"
	"github.com/lib/pq"
	"math/rand"
)

type TransportRepo struct {
	DB          *sql.DB
	TrafficJams []string
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (t *TransportRepo) CreateTrafficjam(report string) error {
	t.TrafficJams = append(t.TrafficJams, report)
	return nil
}

func (t *TransportRepo) GetBusSchedule(busNumber int32) ([]string, error) {
	schedule := []string{}
	err := t.DB.QueryRow("SELECT stations FROM bus WHERE bus_number = $1", busNumber).Scan(pq.Array(&schedule))
	return schedule, err
}

func (t *TransportRepo) TrackBusLocation(busNumber int32) (string, error) {
	schedule := []string{}
	err := t.DB.QueryRow("SELECT stations FROM bus WHERE bus_number = $1", busNumber).Scan(pq.Array(&schedule))

	return schedule[rand.Intn(len(schedule))], err
}
