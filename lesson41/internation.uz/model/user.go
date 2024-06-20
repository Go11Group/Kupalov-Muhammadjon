package model

import "time"

type User struct {
	UserId   string    `json:"user_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password"`
	Time
}

type UserFilter struct {
	Name     *string
	Email    *string
	Birthday *time.Time
	AgeFrom  *int
	AgeTo    *int
	Limit    *int
	Offset   *int
}
