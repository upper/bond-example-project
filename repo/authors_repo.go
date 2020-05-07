package repo

import (
	"github.com/upper/bond-example-project/internal/model"

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

func (authors *AuthorsRepo) FindByID(id uint64) (*Author, error) {
	var author model.Author
	if err := authors.Find(db.Cond{"id": id}).One(&author); err != nil {
		return nil, err
	}
	return NewAuthor(&author), nil
}
