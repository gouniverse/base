package database

import (
	"database/sql"
	"errors"
)

func Execute(ctx QueryableContext, sqlStr string, args ...any) (sql.Result, error) {
	if ctx.queryable == nil {
		return nil, errors.New("querier (db/tx/conn) is nil")
	}

	return ctx.queryable.ExecContext(ctx, sqlStr, args...)
}
