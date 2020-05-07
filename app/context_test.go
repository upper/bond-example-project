package app

import (
	"github.com/upper/bond-example-project/config"
	"github.com/upper/db/adapter/postgresql"
	"github.com/upper/db/bond"

	"context"
)

func testContext() (context.Context, error) {
	sess, err := bond.Open("postgresql",
		postgresql.ConnectionURL{
			Host:     config.Config.Database.Host,
			User:     config.Config.Database.User,
			Database: config.Config.Database.Name,
			Password: config.Config.Database.Password,
		},
	)
	if err != nil {
		return nil, err
	}
	return context.WithValue(
		context.Background(),
		ContextDatabaseSession,
		sess,
	), nil
}
