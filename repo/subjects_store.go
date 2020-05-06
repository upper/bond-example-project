package repo

import (
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
