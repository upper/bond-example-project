package app

import (
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/repo"
	"github.com/upper/db/bond"

	"context"
)

type BooksApp struct {
	repo *repo.BooksRepo
	ctx  context.Context
}

func Books(ctx context.Context) *BooksApp {
	return &BooksApp{
		repo: repo.Books(ctx.Value(ContextDatabaseSession).(bond.Session)),
		ctx:  ctx,
	}
}

func (b *BooksApp) Get(id uint64) (*repo.Book, error) {
	return b.repo.FindByID(id)
}

func (b *BooksApp) Create(post *model.Book) (*repo.Book, error) {
	newBook := repo.NewBook(&model.Book{
		Title:     post.Title,
		AuthorID:  post.AuthorID,
		SubjectID: post.SubjectID,
	})
	if err := b.repo.Create(newBook); err != nil {
		return nil, err
	}
	return newBook, nil
}

func (b *BooksApp) Update(current *model.Book, patch *model.Book) (*repo.Book, error) {
	updatedBook := repo.NewBook(&model.Book{
		ID: current.ID,
	})

	if patch.Title != "" {
		updatedBook.Title = patch.Title
	}
	if patch.AuthorID > 0 {
		updatedBook.AuthorID = patch.AuthorID
	}
	if patch.SubjectID > 0 {
		updatedBook.SubjectID = patch.SubjectID
	}

	if err := b.repo.Update(updatedBook); err != nil {
		return nil, err
	}
	return updatedBook, nil
}

func (b *BooksApp) Delete(book *model.Book) error {
	return b.repo.Delete(repo.NewBook(book))
}

func (b *BooksApp) Paginate(conds ...interface{}) (*Page, error) {
	return PaginateQuery(b.ctx, b.repo.Find(conds), []*model.Book{})
}
