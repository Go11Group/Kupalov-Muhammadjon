package models

import "time"

type Stock struct {
	CompanyName string
	Price       int
	CurrentTime time.Time
	MaxPrice    int
	MinPrice    int
}
