package models

import "time"

type Transaction struct {
	ID        int       `db:"id"`
	UserID    int64     `db:"user_id"`
	Type      string    `db:"type"`
	Category  string    `db:"category"`
	Amount    float64   `db:"amount"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
}

type User struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

type Goal struct {
	ID           int       `db:"id"`
	UserID       int64     `db:"user_id"`
	Title        string    `db:"title"`
	TargetAmount float64   `db:"target_amount"`
	SavedAmount  float64   `db:"saved_amount"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
