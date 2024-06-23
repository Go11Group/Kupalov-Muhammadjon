package models

type Terminal struct {
	Id        string `json:"id"`
	StationId string `json:"station_id"`
}

type CreateTerminal struct {
	StationId string `json:"station_id"`
}

type TerminalFilter struct {
	StationId *string 
}
