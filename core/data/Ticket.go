package data

import "time"

type Ticket struct {
	ID string `db:"id"`
	UserID string `db:"user_id"`
	OrderID string `db:"order_id"`
	Status string `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
