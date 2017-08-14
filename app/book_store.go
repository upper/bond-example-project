package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type BookStore struct {
	bond.Store
}

func Books(sess bond.Session) *BookStore {
	return &BookStore{Store: repo.Authors(sess)}
}

func (bf *BookStore) FindByID(id uint64) db.Result {
	return bf.Find(db.Cond{"id": id})
}
