package dto

import "time"

type Customer struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Surnames string    `json:"surnames"`
	Dob      time.Time `json:"dob"`
}
