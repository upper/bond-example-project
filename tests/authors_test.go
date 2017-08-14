package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
)

func testAuthors(sess bond.Session, t *testing.T) {
	// Attempt to add an author.
	{
		author := app.NewAuthor(&model.Author{
			FirstName: "Adolf",
		})

		err := sess.Save(author)
		assert.Error(t, err, "Expecting error (author has no last name)")
	}
	// Add an author.
	{
		author := app.NewAuthor(&model.Author{
			FirstName: "Adolf",
			LastName:  "Huxley",
		})

		err := sess.Save(author)
		assert.NoError(t, err)
	}
	// Retrive author and fix their name.
	{
		var authors []*app.Author
		err := app.Authors(sess).FindByLastName("Huxley").All(&authors)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(authors))

		aldousHuxley := authors[0]
		assert.Equal(t, "Adolf", aldousHuxley.FirstName)

		aldousHuxley.FirstName = "Aldous"

		err = sess.Save(aldousHuxley)
		assert.NoError(t, err)
	}
	// Make sure the name was changed.
	{
		var aldousHuxley app.Author
		err := app.Authors(sess).FindByLastName("Huxley").One(&aldousHuxley)
		assert.NoError(t, err)

		assert.Equal(t, "Aldous", aldousHuxley.FirstName)
	}
	// Insert a different entry with the same name
	{
		author := app.NewAuthor(&model.Author{
			FirstName: "Aldous",
			LastName:  "Huxley",
		})

		err := sess.Save(author)
		assert.Error(t, err, "Should not be able to add another entry with the same name")
	}
}

func createAuthorsFixtures(sess bond.Session, t *testing.T) {
	authors := []model.Author{
		{
			FirstName: "Aldous",
			LastName:  "Huxley",
		},
		{
			FirstName: "Harper",
			LastName:  "Lee",
		},
		{
			FirstName: "Isaac",
			LastName:  "Asimov",
		},
		{
			FirstName: "Joanne",
			LastName:  "Rowling",
		},
		{
			FirstName: "Orson",
			LastName:  "Welles",
		},
		{
			FirstName: "Virginia",
			LastName:  "Woolf",
		},
	}

	for _, authorData := range authors {
		author := app.NewAuthor(&authorData)
		err := sess.Save(author)
		assert.NoError(t, err)
	}
}

func TestAuthors(t *testing.T) {
	{
		err := repo.Session.SessionTx(nil, func(tx bond.Session) error {
			cleanUp(tx, t)
			testAuthors(tx, t)
			return nil
		})
		assert.NoError(t, err)
	}

	{
		err := repo.Session.SessionTx(nil, func(tx bond.Session) error {
			cleanUp(tx, t)
			createAuthorsFixtures(tx, t)
			return nil
		})
		assert.NoError(t, err)
	}
}
