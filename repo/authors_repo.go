package repo

import (
	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type AuthorsRepo struct {
	bond.Store
}

func Authors(sess bond.Session) *AuthorsRepo {
	return &AuthorsRepo{
		Store: sess.Store("authors"),
	}
}

func (authors *AuthorsRepo) FindByLastName(lastName string) db.Result {
	return authors.Find(db.Cond{"last_name": lastName})
}

func (authors *AuthorsRepo) FindByID(id uint64) db.Result {
	return authors.Find(db.Cond{"id": id})
}
