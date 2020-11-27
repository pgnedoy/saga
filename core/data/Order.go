package data

import "time"

type Order struct {
	ID string
	Name string
	Quantity int
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}