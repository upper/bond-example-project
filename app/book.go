package app

import (
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"errors"

	"upper.io/bond"
	"upper.io/db.v3"
)

type Book struct {
	*model.Book `json:"book"`
}

func (s *Book) Store(sess bond.Session) bond.Store {
	return repo.Books(sess)
}

func NewBook(book *model.Book) *Book {
	return &Book{book}
}

func (book *Book) Validate() error {
	if book.AuthorID == 0 {
		return errors.New("Missing author")
	}
	if book.SubjectID == 0 {
		return errors.New("Missing subject")
	}
	if book.Title == "" {
		return errors.New("Missing title")
	}
	return nil
}

func (b *Book) BeforeCreate(sess bond.Session) error {
	{
		conds := db.Cond{"title": b.Title, "author_id": b.AuthorID}
		if b.ID != 0 {
			conds["id"] = b.ID
		}

		exists, err := repo.Books(sess).Find(conds).Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A different entry with the same title and author already exists")
		}
	}
	return nil
}

var _ interface {
	bond.Model
	bond.HasBeforeCreate
} = &Book{}
