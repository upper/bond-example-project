package repo

import (
	"github.com/upper/bond-example-project/internal/model"

	"errors"

	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type Book struct {
	*model.Book `json:"book"`
}

var _ interface {
	bond.Model
	bond.HasBeforeCreate
} = &Book{}

func (b *Book) Store(sess bond.Session) bond.Store {
	return Books(sess)
}

func NewBook(book *model.Book) *Book {
	return &Book{
		Book: book,
	}
}

func (b *Book) Validate() error {
	if b.AuthorID == 0 {
		return errors.New("Missing author")
	}
	if b.SubjectID == 0 {
		return errors.New("Missing subject")
	}
	if b.Title == "" {
		return errors.New("Missing title")
	}
	return nil
}

func (b *Book) BeforeSave(sess bond.Session) error {
	return nil
}

func (b *Book) BeforeCreate(sess bond.Session) error {
	{
		conds := db.Cond{"title": b.Title, "author_id": b.AuthorID}
		if b.ID != 0 {
			conds["id"] = b.ID
		}

		exists, err := Books(sess).Find(conds).Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A different entry with the same title and author already exists")
		}
	}
	return nil
}
