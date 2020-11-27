package data

import "time"

type Ticket struct {
	ID string
	UserID string
	OrderID string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}
