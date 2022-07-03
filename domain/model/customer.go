package model

import "time"

type Customer struct {
	ID       int
	Name     string
	Surnames string
	Dob      time.Time
}
