package app

import (
	"context"

	"github.com/upper/db/bond"
)

const (
	ContextDatabaseSession = "app.database.session"

	ContextPaginationNext   = "app.pagination.next"
	ContextPaginationPrev   = "app.pagination.prev"
	ContextPaginationPage   = "app.pagination.page"
	ContextPaginationCursor = "app.pagination.cursor"
	ContextPaginationSize   = "app.pagination.size"
)

func WithDatabaseSession(ctx context.Context, sess bond.Session) context.Context {
	return context.WithValue(ctx, ContextDatabaseSession, sess)
}
