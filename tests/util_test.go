package tests

import (
	"fmt"
	"github.com/upper/bond-example-project/repo"
	"testing"

	"upper.io/bond"
)

func cleanUp(sess bond.Session, t *testing.T) {
	stores := []func(sess bond.Session) bond.Store{
		repo.Authors,
		repo.Subjects,
		repo.Books,
	}

	for _, store := range stores {
		tableName := store(sess).Name()
		_, err := sess.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName))
		if err != nil {
			t.Fatalf("Failed to truncate table %v: %v", tableName, err)
		}
	}
}
