package data

import "time"

type Consumer struct {
	ID string `db:"id"`
	FirstName string `db:"first_name"`
	SecondName string `db:"second_name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
