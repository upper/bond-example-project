package model

import (
	"time"
)

// Subject represents a book's subject.
type Subject struct {
	ID       uint64 `db:"id,omitempty" json:"id"`
	Name     string `db:"name" json:"name"`
	Location string `db:"location" json:"location"`

	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}
