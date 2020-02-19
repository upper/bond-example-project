package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upper/bond-example-project/app"
	"github.com/upper/bond-example-project/model"
	"github.com/upper/bond-example-project/repo"

	"upper.io/bond"
)

func testSubjects(sess bond.Session, t *testing.T) {
	// Attempt to add a subject.
	{
		subject := app.NewSubject(&model.Subject{
			Location: "Evergreen St.",
		})

		err := sess.Save(subject)
		assert.Error(t, err, "Expecting error (subject has no name)")
	}
	// Add an subject.
	{
		subject := app.NewSubject(&model.Subject{
			Location: "Evergreen St.",
			Name:     "Art",
		})

		err := sess.Save(subject)
		assert.NoError(t, err, "Should be able to add a new subject")

		assert.NotZero(t, subject.ID)

		// Save again without changing anything.
		err = sess.Save(subject)
		assert.NoError(t, err, "Should be able to save a subject")

		assert.NotZero(t, subject.ID)

		// Change name and save.
		subject.Name = "Arts"
		err = sess.Save(subject)
		assert.NoError(t, err, "Should be able to save a subject")
	}
	// Attempt to add another subject with the same name.
	{
		subject := app.NewSubject(&model.Subject{
			Location: "Evergreen St.",
			Name:     "Arts",
		})

		err := sess.Save(subject)
		assert.Error(t, err, "Should not be able to add a second subject with a name that already exists")
	}
}

func createSubjectsFixtures(sess bond.Session, t *testing.T) {
	subjects := []model.Subject{
		{
			Name:     "Arts",
			Location: "Creativity St",
		},
		{
			Name:     "Science Fiction",
			Location: "Academic Rd",
		},
		{
			Name:     "Classics",
			Location: "Main St",
		},
		{
			Name:     "Entertainment",
			Location: "Main St",
		},
		{
			Name:     "Business",
			Location: "Black Raven Dr",
		},
		{
			Name:     "Mistery",
			Location: "Black Raven Dr",
		},
		{
			Name:     "Drama",
			Location: "Academic Rd",
		},
		{
			Name:     "Computers",
			Location: "Creativity St",
		},
		{
			Name:     "Children",
			Location: "Productivity Ave",
		},
	}

	for _, subjectData := range subjects {
		subject := app.NewSubject(&subjectData)
		err := sess.Save(subject)
		assert.NoError(t, err)
	}
}

func TestSubjects(t *testing.T) {
	{
		err := repo.Session.SessionTx(nil, func(tx bond.Session) error {
			cleanUp(tx, t)
			testSubjects(tx, t)
			return nil
		})
		assert.NoError(t, err)
	}

	{
		err := repo.Session.SessionTx(nil, func(tx bond.Session) error {
			cleanUp(tx, t)
			createSubjectsFixtures(tx, t)
			return nil
		})
		assert.NoError(t, err)
	}
}
