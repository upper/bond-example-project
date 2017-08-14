package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type AuthorStore struct {
	bond.Store
}

func Authors(sess bond.Session) *AuthorStore {
	return &AuthorStore{Store: repo.Authors(sess)}
}

func (af *AuthorStore) FindByLastName(lastName string) db.Result {
	return af.Find(db.Cond{"last_name": lastName})
}

func (af *AuthorStore) FindByID(id uint64) db.Result {
	return af.Find(db.Cond{"id": id})
}
