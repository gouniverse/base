package database

import (
	"database/sql"
	"errors"
)

func Query(ctx QueryableContext, sqlStr string, args ...any) (*sql.Rows, error) {
	if ctx.queryable == nil {
		return nil, errors.New("querier (db/tx/conn) is nil")
	}

	return ctx.queryable.QueryContext(ctx, sqlStr, args...)
}
