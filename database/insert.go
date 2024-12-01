package database

import (
	"database/sql"
	"errors"
)

// func InsertOne(ctx QueryableContext, tableName string, data map[string]any) (int64, error) {
// 	return 0, errors.New("not implemented")
// }

// func InsertMany(ctx QueryableContext, tableName string, data []map[string]any) (int64, error) {
// 	return 0, errors.New("not implemented")
// }

func Execute(ctx QueryableContext, sqlStr string, args ...any) (sql.Result, error) {
	if ctx.queryable == nil {
		return nil, errors.New("querier (db/tx/conn) is nil")
	}

	return ctx.queryable.ExecContext(ctx, sqlStr, args...)
}

func Query(ctx QueryableContext, sqlStr string, args ...any) (*sql.Rows, error) {
	if ctx.queryable == nil {
		return nil, errors.New("querier (db/tx/conn) is nil")
	}

	return ctx.queryable.QueryContext(ctx, sqlStr, args...)
}
