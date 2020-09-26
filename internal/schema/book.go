package schema

import (
	"time"
)

// Book represents a book.
type Book struct {
	ID    uint64 `db:"id,omitempty" json:"id"`
	Title string `db:"title" json:"title"`
	ISBN  string `db:"isbn" json:"isbn"`

	PublisherID uint64     `db:"publisher_id" json:"publisher_id"`
	Publisher   *Publisher `db:"-" json:"publisher"`

	Category Category `db:"category" json:"category"`

	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}
