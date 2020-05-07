package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/upper/bond-example-project/internal/model"
	"testing"

	"context"
)

type AuthorsTestSuite struct {
	suite.Suite

	app *AuthorsApp
	ctx context.Context
}

func (suite *AuthorsTestSuite) SetupTest() {
	var err error
	suite.ctx, err = testContext()
	assert.NoError(suite.T(), err)

	suite.app = Authors(suite.ctx)

	err = Books(suite.ctx).repo.Find().Delete()
	assert.NoError(suite.T(), err)

	err = suite.app.repo.Find().Delete()
	assert.NoError(suite.T(), err)
}

func (suite *AuthorsTestSuite) TestCreateAuthor() {
	post := model.Author{
		FirstName: "Harper",
		LastName:  "Lee",
	}
	newAuthor, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), newAuthor)

	assert.Equal(suite.T(), newAuthor.FirstName, post.FirstName)
	assert.Equal(suite.T(), newAuthor.LastName, post.LastName)
}

func (suite *AuthorsTestSuite) TestGetAuthor() {
	post := model.Author{
		FirstName: "George",
		LastName:  "Orwell",
	}
	newAuthor, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	storedAuthor, err := suite.app.Get(newAuthor.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), storedAuthor)

	assert.Equal(suite.T(), storedAuthor.FirstName, newAuthor.FirstName)
	assert.Equal(suite.T(), storedAuthor.LastName, newAuthor.LastName)
}

func (suite *AuthorsTestSuite) TestUpdateAuthor() {
	post := model.Author{
		FirstName: "Isaa",
		LastName:  "Asimov",
	}
	newAuthor, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	updatedAuthor, err := suite.app.Update(newAuthor.Author, &model.Author{
		FirstName: "Isaac",
	})
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), updatedAuthor)

	assert.Equal(suite.T(), "Isaac", updatedAuthor.FirstName)
	assert.Equal(suite.T(), "Asimov", updatedAuthor.LastName)
}

func (suite *AuthorsTestSuite) TestDeleteAuthor() {
	post := model.Author{
		FirstName: "Pepe",
		LastName:  "Pecas",
	}
	newAuthor, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	err = suite.app.Delete(newAuthor.Author)
	assert.NoError(suite.T(), err)

	_, err = suite.app.Get(newAuthor.ID)
	assert.Error(suite.T(), err)
}

func TestAuthors(t *testing.T) {
	suite.Run(t, new(AuthorsTestSuite))
}
