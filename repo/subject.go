package repo

import (
	"github.com/upper/bond-example-project/internal/model"

	"errors"

	"github.com/upper/db"
	"github.com/upper/db/bond"
)

type Subject struct {
	*model.Subject `json:"subject"`
}

var _ interface {
	bond.Model
	bond.BeforeCreateHook
} = &Subject{}

func (s *Subject) Store(sess bond.Session) bond.Store {
	return Subjects(sess)
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

		exists, err := s.Store(sess).Find(conds).Exists()
		if err != nil {
			return err
		}
		if exists {
			return errors.New("A different entry with the same name already exists")
		}
	}
	return nil
}
