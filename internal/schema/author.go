package schema

import "time"

// Author represents the author of a book
type Author struct {
	ID      uint64 `db:"id,omitempty" json:"-"`
	Name    string `db:"name,omitempty" json:"name"`
	Surname string `db:"surname,omitempty" json:"surname"`

	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}
