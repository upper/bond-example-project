package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
)

func testBooks(sess bond.Session, t *testing.T) {
	{
		book := app.NewBook(&model.Book{
			Title: "Brave New World",
		})

		err := sess.Save(book)
		assert.Error(t, err, "Should not have been able to save without subject nor author")
	}
	{
		var author app.Author
		err := app.Authors(sess).FindByLastName("Huxley").One(&author)
		assert.NoError(t, err)
		assert.NotNil(t, author, "Expecting an author")

		book := app.NewBook(&model.Book{
			Title:    "Brave New World",
			AuthorID: author.ID,
		})

		err = sess.Save(book)
		assert.Error(t, err, "Should not have been able to save without a subject")
	}
	{
		var author app.Author
		err := app.Authors(sess).FindByLastName("Huxley").One(&author)
		assert.NoError(t, err)
		assert.NotZero(t, author.ID)

		var subject app.Subject
		err = app.Subjects(sess).FindByName("Science Fiction").One(&subject)
		assert.NoError(t, err)
		assert.NotZero(t, subject.ID)

		book := app.NewBook(&model.Book{
			Title:     "Brave New World",
			AuthorID:  author.ID,
			SubjectID: subject.ID,
		})

		err = sess.Save(book)
		assert.NoError(t, err)

		err = sess.Save(book)
		assert.NoError(t, err)
	}
	{
		var author app.Author
		err := app.Authors(sess).FindByLastName("Huxley").One(&author)

		assert.NoError(t, err)
		assert.NotNil(t, author)
		assert.NotZero(t, author.ID)

		var subject app.Subject
		err = app.Subjects(sess).FindByName("Science Fiction").One(&subject)
		assert.NoError(t, err)
		assert.NotZero(t, subject.ID)

		book := app.NewBook(&model.Book{
			Title:     "Brave New World",
			AuthorID:  author.ID,
			SubjectID: subject.ID,
		})

		err = sess.Save(book)
		assert.Error(t, err, "Should not be able to add a duplicated record")
	}
}

func TestBooks(t *testing.T) {
	err := repo.Session.SessionTx(nil, func(tx bond.Session) error {

		cleanUp(tx, t)

		createSubjectsFixtures(tx, t)
		createAuthorsFixtures(tx, t)

		testBooks(tx, t)

		return nil
	})
	assert.NoError(t, err)
}
