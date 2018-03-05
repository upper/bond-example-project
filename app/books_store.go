package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type BooksStore struct {
	bond.Store
}

func Books(sess bond.Session) *BooksStore {
	return &BooksStore{Store: repo.Books(sess)}
}

func (s *BooksStore) FindByID(id uint64) db.Result {
	return s.Find(db.Cond{"id": id})
}
