package app

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/upper/bond-example-project/internal/model"
	"testing"

	"context"
)

type SubjectsTestSuite struct {
	suite.Suite

	app *SubjectsApp
	ctx context.Context
}

func (suite *SubjectsTestSuite) SetupTest() {
	var err error
	suite.ctx, err = testContext()
	assert.NoError(suite.T(), err)

	err = Books(suite.ctx).repo.Find().Delete()
	assert.NoError(suite.T(), err)

	suite.app = Subjects(suite.ctx)
	err = suite.app.repo.Find().Delete()
	assert.NoError(suite.T(), err)
}

func (suite *SubjectsTestSuite) TestCreateSubject() {
	post := model.Subject{
		Name:     "Fiction",
		Location: "Aisle A",
	}
	newSubject, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), newSubject)

	assert.Equal(suite.T(), newSubject.Name, post.Name)
	assert.Equal(suite.T(), newSubject.Location, post.Location)
}

func (suite *SubjectsTestSuite) TestGetSubject() {
	post := model.Subject{
		Name:     "Science",
		Location: "Aisle B",
	}
	newSubject, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	storedSubject, err := suite.app.Get(newSubject.ID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), storedSubject)

	assert.Equal(suite.T(), storedSubject.Name, newSubject.Name)
	assert.Equal(suite.T(), storedSubject.Location, newSubject.Location)
}

func (suite *SubjectsTestSuite) TestUpdateSubject() {
	post := model.Subject{
		Name:     "Sciec",
		Location: "Aisle C",
	}
	newSubject, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	updatedSubject, err := suite.app.Update(newSubject.Subject, &model.Subject{
		Name: "Science",
	})
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), updatedSubject)

	assert.Equal(suite.T(), "Science", updatedSubject.Name)
}

func (suite *SubjectsTestSuite) TestDeleteSubject() {
	post := model.Subject{
		Name:     "Fantasy",
		Location: "Aisle D",
	}
	newSubject, err := suite.app.Create(&post)
	assert.NoError(suite.T(), err)

	err = suite.app.Delete(newSubject.Subject)
	assert.NoError(suite.T(), err)

	_, err = suite.app.Get(newSubject.ID)
	assert.Error(suite.T(), err)
}

func TestSubjects(t *testing.T) {
	suite.Run(t, new(SubjectsTestSuite))
}
