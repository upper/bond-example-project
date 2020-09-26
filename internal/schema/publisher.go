package schema

// Publisher represents the publisher of a
type Publisher struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
