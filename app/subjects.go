package app

import (
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/repo"
	"github.com/upper/db/bond"

	"context"
)

type SubjectsApp struct {
	repo *repo.SubjectsRepo
	ctx  context.Context
}

func Subjects(ctx context.Context) *SubjectsApp {
	return &SubjectsApp{
		repo: repo.Subjects(ctx.Value(ContextDatabaseSession).(bond.Session)),
		ctx:  ctx,
	}
}

func (b *SubjectsApp) Get(id uint64) (*repo.Subject, error) {
	return b.repo.FindByID(id)
}

func (b *SubjectsApp) Create(post *model.Subject) (*repo.Subject, error) {
	newSubject := repo.NewSubject(&model.Subject{
		Name:     post.Name,
		Location: post.Location,
	})
	if err := b.repo.Create(newSubject); err != nil {
		return nil, err
	}
	return newSubject, nil
}

func (b *SubjectsApp) Update(current *model.Subject, patch *model.Subject) (*repo.Subject, error) {
	updatedSubject := repo.NewSubject(&model.Subject{
		ID: current.ID,
	})

	if patch.Name != "" {
		updatedSubject.Name = patch.Name
	}
	if patch.Location != "" {
		updatedSubject.Location = patch.Location
	}

	if err := b.repo.Update(updatedSubject); err != nil {
		return nil, err
	}
	return updatedSubject, nil
}

func (b *SubjectsApp) Delete(book *model.Subject) error {
	return b.repo.Delete(repo.NewSubject(book))
}

func (b *SubjectsApp) Paginate(conds ...interface{}) (*Page, error) {
	return PaginateQuery(b.ctx, b.repo.Find(conds), []*model.Subject{})
}
