package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/upper/bond-example-project/internal/model"
	"github.com/upper/bond-example-project/repo"
	"testing"

	"context"
)

type BooksTestSuite struct {
	suite.Suite
	app *BooksApp
	ctx context.Context

	fixtures struct {
		subject *repo.Subject
		author  *repo.Author
	}
}

func (suite *BooksTestSuite) SetupTest() {
	var err error
	suite.ctx, err = testContext()
	assert.NoError(suite.T(), err)

	suite.app = Books(suite.ctx)
	err = suite.app.repo.Find().Delete()
	assert.NoError(suite.T(), err)

	err = Authors(suite.ctx).repo.Find().Delete()
	assert.NoError(suite.T(), err)
	suite.fixtures.author, err = Authors(suite.ctx).Create(&model.Author{
		FirstName: "Jane",
		LastName:  "Doe",
	})
	assert.NoError(suite.T(), err)

	err = Subjects(suite.ctx).repo.Find().Delete()
	assert.NoError(suite.T(), err)
	suite.fixtures.subject, err = Subjects(suite.ctx).Create(&model.Subject{
		Name:     "General",
		Location: "Aisle A",
	})
	assert.NoError(suite.T(), err)
}

func (suite *BooksTestSuite) TestCreateBook() {
	post := model.Book{
		Title:     "The Sun Also Rises",
		SubjectID: suite.fixtures.subject.ID,
		AuthorID:  suite.fixtures.author.ID,
	}
	newBook, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), newBook)

	assert.Equal(suite.T(), newBook.Title, post.Title)
}

func (suite *BooksTestSuite) xTestGetBook() {
	post := model.Book{
		Title:     "1984",
		SubjectID: suite.fixtures.subject.ID,
		AuthorID:  suite.fixtures.author.ID,
	}
	newBook, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	storedBook, err := suite.app.Get(newBook.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), storedBook)

	assert.Equal(suite.T(), storedBook.ID, newBook.ID)
	assert.Equal(suite.T(), storedBook.Title, newBook.Title)
}

func (suite *BooksTestSuite) xTestUpdateBook() {
	post := model.Book{
		Title: "How to Kill a Mockingbird",
	}
	newBook, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), newBook)

	assert.Equal(suite.T(), newBook.Title, post.Title)

	storedBook, err := suite.app.Get(newBook.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), storedBook)

	assert.Equal(suite.T(), storedBook.ID, newBook.ID)
	assert.Equal(suite.T(), storedBook.Title, newBook.Title)
}

func TestBooks(t *testing.T) {
	suite.Run(t, new(BooksTestSuite))
}
