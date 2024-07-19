package models

type Item struct {
	Title       string
	Description string
	Price       float64
}

type ItemInfo struct {
	Id          string
	Title       string
	Description string
	Price       float64
	Time
}

type ItemUpdate struct {
	Id          string
	Title       string
	Description string
	Price       float64
}
