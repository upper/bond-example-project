package model

import (
	"time"
)

// Book represents a book.
type Book struct {
	ID        uint64 `db:"id,omitempty" json:"id"`
	Title     string `db:"title" json:"title"`
	AuthorID  uint64 `db:"author_id" json:"author_id"`
	SubjectID uint64 `db:"subject_id" json:"subject_id"`

	UpdatedAt time.Time `db:"updated_at,omitempty" json:"updated_at"`
	CreatedAt time.Time `db:"created_at,omitempty" json:"created_at"`
}
