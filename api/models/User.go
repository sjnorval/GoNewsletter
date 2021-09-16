package models

type User struct {
	ID          string  `json:"id"`
	Name        string  `validate:"required"`
	Surname     string  `validate:"required"`
	Email       string  `validate:"required,email"`
	LastUpdated string  `json:"lastupdated"`
	Topics      []Topic `json: "topics"`
}
