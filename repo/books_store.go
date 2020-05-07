package repo

import (
	"github.com/upper/bond-example-project/internal/model"

	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type BooksRepo struct {
	bond.Store
}

func Books(sess bond.Session) *BooksRepo {
	return &BooksRepo{
		Store: sess.Store("books"),
	}
}

func (books *BooksRepo) FindByID(id uint64) (*Book, error) {
	var book model.Book
	if err := books.Find(db.Cond{"id": id}).One(&book); err != nil {
		return nil, err
	}
	return NewBook(&book), nil
}
