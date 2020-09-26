package schema

import (
	"time"
)

type Review struct {
	ID     int64 `db:"id,omitempty" json:"-"`
	BookID int64 `db:"book_id" json:"book_id"`

	ReviewerID int64  `db:"reviewer_id" json:"-"`
	Reviewer   string `db:"-" json:"reviewer"`

	Content string `db:"content" json:"content"`

	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
}
