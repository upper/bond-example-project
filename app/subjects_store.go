package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type SubjectsStore struct {
	bond.Store
}

func Subjects(sess bond.Session) *SubjectsStore {
	return &SubjectsStore{Store: repo.Subjects(sess)}
}

func (s *SubjectsStore) FindByName(name string) db.Result {
	return s.Find(db.Cond{"name": name})
}
