package repo

import (
	"github.com/upper/bond-example-project/internal/model"

	"context"
	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type BooksRepo struct {
	bond.Store
	ctx context.Context
}

func Books(sess bond.Session) *BooksRepo {
	return &BooksRepo{
		Store: sess.Store("books"),
		ctx:   context.Background(),
	}
}

func (books *BooksRepo) FindByID(id uint64) (*Book, error) {
	var book model.Book
	if err := books.Find(db.Cond{"id": id}).One(&book); err != nil {
		return nil, err
	}
	return NewBook(&book), nil
}

func (books *BooksRepo) Paginate(conds ...interface{}) db.Result {
	/*
		operation := func(sess bond.Session) error {
			var books []*app.Book
			var err error

			page, err = ws.Paginate(r, app.Books(sess).Find(), &books)
			if err != nil {
				return err
			}

			return nil
		}
	*/
	return books.Find(conds)
}
