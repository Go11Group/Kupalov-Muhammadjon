package models

type Station struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateStation struct {
	Name string `json:"name"`
}

type StationFilter struct {
	Name *string
}
