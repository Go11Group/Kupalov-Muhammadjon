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
	Limit    *int
	Offset   *int
}
