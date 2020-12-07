package data

import "time"

type Order struct {
	ID string `db:"id"`
	Name string `db:"name"`
	ConsumerID string `db:"consumer_id"`
	Quantity int `db:"quantity"`
	Status string `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}