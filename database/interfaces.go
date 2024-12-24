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

	// PrepareContext creates a prepared statement for use within a transaction.
	//
	// The returned statement operates within the transaction and will be closed
	// when the transaction has been committed or rolled back.
	//
	// To use an existing prepared statement on this transaction, see [Tx.Stmt].
	//
	// The provided context will be used for the preparation of the context, not
	// for the execution of the returned statement. The returned statement
	// will run in the transaction context.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

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

// Verify that sql.DB, sql.Conn and sql.Tx implement the QueryableInterface
var (
	_ QueryableInterface = &sql.DB{}
	_ QueryableInterface = &sql.Conn{}
	_ QueryableInterface = &sql.Tx{}
)
