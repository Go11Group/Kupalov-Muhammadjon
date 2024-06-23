package models

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type CreateUpdateUser struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type UserFilter struct {
	Name    *string `json:"name,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	AgeFrom *int    `json:"age_from,omitempty"`
	AgeTo   *int    `json:"age_to,omitempty"`
}
