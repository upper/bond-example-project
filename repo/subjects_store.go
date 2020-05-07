package repo

import (
	"github.com/upper/bond-example-project/internal/model"

	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type SubjectsRepo struct {
	bond.Store
}

func Subjects(sess bond.Session) *SubjectsRepo {
	return &SubjectsRepo{
		Store: sess.Store("subjects"),
	}
}

func (subjects *SubjectsRepo) FindByName(name string) db.Result {
	return subjects.Find(db.Cond{"name": name})
}

func (subjects *SubjectsRepo) FindByID(id uint64) (*Subject, error) {
	var subject model.Subject
	if err := subjects.Find(db.Cond{"id": id}).One(&subject); err != nil {
		return nil, err
	}
	return NewSubject(&subject), nil
}
