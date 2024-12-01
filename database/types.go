package database

import (
	"context"
	"database/sql"
)

// Queryable is an interface that defines a set of methods for executing
// SQL queries on a database or data source.
//
// It can be one of the following:
// - *sql.DB
// - *sql.Conn
// - *sql.Tx
//
// Implementations of this interface provide a way to execute queries in a context,
// allowing for cancellation and timeout control.
type QueryableInterface interface {
	// ExecContext executes a SQL query in the given context.
	// It returns a sql.Result object containing information about the execution,
	// or an error if the query failed.
	//
	// The context is used to control the execution of the query, allowing for
	// cancellation and timeout control.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// QueryContext executes a SQL query in the given context and returns a
	// *sql.Rows object containing the query results.
	//
	// The context is used to control the execution of the query, allowing for
	// cancellation and timeout control.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowContext executes a SQL query in the given context and returns a
	// *sql.Row object containing a single row of results.
	//
	// The context is used to control the execution of the query, allowing for
	// cancellation and timeout control.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

var (
	_ QueryableInterface = &sql.DB{}
	_ QueryableInterface = &sql.Conn{}
	_ QueryableInterface = &sql.Tx{}
)

// Context returns a new context with the given QueryableInterface.
func Context(ctx context.Context, queryable QueryableInterface) QueryableContext {
	return NewQueryableContext(ctx, queryable)
}

// NewQueryableContext returns a new context with the given QueryableInterface.
func NewQueryableContext(ctx context.Context, queryable QueryableInterface) QueryableContext {
	return QueryableContext{Context: ctx, queryable: queryable}
}

// TransactionContext is a context that contains a transaction.
type QueryableContext struct {
	context.Context
	queryable QueryableInterface
}

var _ context.Context = QueryableContext{}

func IsQueryableContext(ctx context.Context) bool {
	return ctx != nil && ctx.Value("queryable") != nil
}
