package schema

type Reviewer struct {
	ID    int64  `db:"id,omitempty" json:"-"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
