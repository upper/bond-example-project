package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type AuthorsStore struct {
	bond.Store
}

func Authors(sess bond.Session) *AuthorsStore {
	return &AuthorsStore{Store: repo.Authors(sess)}
}

func (s *AuthorsStore) FindByLastName(lastName string) db.Result {
	return s.Find(db.Cond{"last_name": lastName})
}

func (s *AuthorsStore) FindByID(id uint64) db.Result {
	return s.Find(db.Cond{"id": id})
}
