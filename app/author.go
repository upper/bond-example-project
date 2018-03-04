package app

import (
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"errors"

	"upper.io/bond"
	"upper.io/db.v3"
)

type Author struct {
	*model.Author `json:"author"`
}

func (a *Author) Store(sess bond.Session) bond.Store {
	return repo.Authors(sess)
}

func NewAuthor(a *model.Author) *Author {
	return &Author{a}
}

func (a *Author) Validate() error {
	if a.FirstName == "" {
		return errors.New("Missing author's first name")
	}
	if a.LastName == "" {
		return errors.New("Missing author's last name")
	}
	return nil
}

func (a *Author) BeforeCreate(sess bond.Session) error {
	{
		conds := db.Cond{"first_name": a.FirstName, "last_name": a.LastName}
		if a.ID != 0 {
			conds["id"] = a.ID
		}

		exists, err := Authors(sess).Find(conds).Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A different entry with the same first name and last name already exists")
		}
	}
	return nil
}

var _ interface {
	bond.Model
	bond.HasBeforeCreate
} = &Author{}
