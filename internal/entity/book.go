package entity

import "time"

// Book represents an book record.
type Book struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
