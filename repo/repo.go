package repo

import (
	"log"
	"os"

	"upper.io/bond"
	"upper.io/db.v3/postgresql"
)

var Session bond.Session

func Books(sess bond.Session) bond.Store {
	return sess.Store("books")
}

func Authors(sess bond.Session) bond.Store {
	return sess.Store("authors")
}

func Subjects(sess bond.Session) bond.Store {
	return sess.Store("subjects")
}

func loadSettings() postgresql.ConnectionURL {
	return postgresql.ConnectionURL{
		Host:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USERNAME"),
		Database: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func init() {
	settings := loadSettings()

	var err error
	Session, err = bond.Open("postgresql", settings)
	if err != nil {
		log.Fatal("postgresql.Open", err)
	}
}
