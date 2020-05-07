package app

import (
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/repo"
	"github.com/upper/db/bond"

	"context"
)

type AuthorsApp struct {
	repo *repo.AuthorsRepo
	ctx  context.Context
}

func Authors(ctx context.Context) *AuthorsApp {
	return &AuthorsApp{
		repo: repo.Authors(ctx.Value(ContextDatabaseSession).(bond.Session)),
		ctx:  ctx,
	}
}

func (b *AuthorsApp) Get(id uint64) (*repo.Author, error) {
	return b.repo.FindByID(id)
}

func (b *AuthorsApp) Create(post *model.Author) (*repo.Author, error) {
	newAuthor := repo.NewAuthor(&model.Author{
		FirstName: post.FirstName,
		LastName:  post.LastName,
	})
	if err := b.repo.Create(newAuthor); err != nil {
		return nil, err
	}
	return newAuthor, nil
}

func (b *AuthorsApp) Update(current *model.Author, patch *model.Author) (*repo.Author, error) {
	updatedAuthor := repo.NewAuthor(&model.Author{
		ID: current.ID,
	})

	if patch.FirstName != "" {
		updatedAuthor.FirstName = patch.FirstName
	}
	if patch.LastName != "" {
		updatedAuthor.LastName = patch.LastName
	}

	if err := b.repo.Update(updatedAuthor); err != nil {
		return nil, err
	}
	return updatedAuthor, nil
}

func (b *AuthorsApp) Delete(author *model.Author) error {
	return b.repo.Delete(repo.NewAuthor(author))
}

func (b *AuthorsApp) Paginate(conds ...interface{}) (*Page, error) {
	return PaginateQuery(b.ctx, b.repo.Find(conds), []*model.Author{})
}
