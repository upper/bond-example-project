package app

import (
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
	"upper.io/db.v3"
)

type SubjectStore struct {
	bond.Store
}

func Subjects(sess bond.Session) *SubjectStore {
	return &SubjectStore{Store: repo.Subjects(sess)}
}

func (sf *SubjectStore) FindByName(name string) db.Result {
	return sf.Find(db.Cond{"name": name})
}
