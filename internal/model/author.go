package model

import "time"

// Author represents the author of a book
type Author struct {
	ID        uint64 `db:"id,omitempty" json:"id"`
	FirstName string `db:"first_name,omitempty" json:"first_name"`
	LastName  string `db:"last_name,omitempty" json:"last_name"`

	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}
