package app

import (
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"errors"
	"upper.io/bond"
	"upper.io/db.v3"
)

type Subject struct {
	*model.Subject `json:"subject"`
}

func (s *Subject) Store(sess bond.Session) bond.Store {
	return repo.Subjects(sess)
}

func NewSubject(subject *model.Subject) *Subject {
	return &Subject{subject}
}

func (s *Subject) Validate() error {
	if s.Name == "" {
		return errors.New("Missing subject's name")
	}
	return nil
}

func (s *Subject) BeforeCreate(sess bond.Session) error {
	{
		conds := db.Cond{"name": s.Name}
		if s.ID != 0 {
			conds["id"] = s.ID
		}

		exists, err := Subjects(sess).Find(conds).Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A different entry with the same name already exists")
		}
	}
	return nil
}

var _ interface {
	bond.Model
	bond.HasBeforeCreate
} = &Subject{}
