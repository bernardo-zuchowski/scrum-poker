package models

import "time"

type Todo struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
