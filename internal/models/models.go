package models

import "time"

type ErrorRes struct {
	Error string `json:"error"`
}

type Request struct {
	Task        string `json:"task" db:"task"`
	Description string `json:"description" db:"description"`
	Do          bool   `json:"done" db:"done"`
}

type GetByIDRequest struct {
	ID int `json:"id" db:"id"`
}

type Task struct {
	ID          int       `json:"id" db:"id"`
	Task        string    `json:"task" db:"task"`
	Description string    `json:"description" db:"description"`
	Do          bool      `json:"done" db:"done"`
	Created_at  time.Time `json:"created_at" db:"created_at"`
	Updated_at  time.Time `json:"updated_at" db:"updated_at"`
}
