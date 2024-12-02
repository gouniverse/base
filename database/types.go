package database

import (
	"context"
	"database/sql"
)

// NewQueryableContext returns a new context with the given QueryableInterface.
func NewQueryableContext(ctx context.Context, queryable QueryableInterface) QueryableContext {
	return QueryableContext{Context: ctx, queryable: queryable}
}

// Verify that QueryableContext implements the context.Context interface.
var _ context.Context = QueryableContext{}

// QueryableContext extends the context.Context interface with a queryable field.
// The queryable field may be of type *sql.DB, *sql.Conn, or *sql.Tx.
type QueryableContext struct {
	context.Context
	queryable QueryableInterface
}

func (ctx QueryableContext) IsDB() bool {
	if ctx.queryable == nil {
		return false
	}

	if ctx.IsTx() {
		return false
	}

	if ctx.IsConn() {
		return false
	}

	_, ok := ctx.queryable.(*sql.DB)

	return ok
}

func (ctx QueryableContext) IsConn() bool {
	if ctx.queryable == nil {
		return false
	}

	_, ok := ctx.queryable.(*sql.Conn)

	return ok
}

func (ctx QueryableContext) IsTx() bool {
	if ctx.queryable == nil {
		return false
	}

	_, ok := ctx.queryable.(*sql.Tx)

	return ok
}

func (ctx QueryableContext) Queryable() QueryableInterface {
	return ctx.queryable
}
