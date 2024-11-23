package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/gouniverse/maputils"
)

// Querier is something that we can query and get the *sql.Rows from.
// For example, it can be: *sql.DB, *sql.Conn or *sql.Tx.
type Querier interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

var (
	_ Querier = &sql.DB{}
	_ Querier = &sql.Conn{}
	_ Querier = &sql.Tx{}
)

func SelectToMapAny(ctx context.Context, q Querier, sqlStr string, args ...any) ([]map[string]any, error) {
	if q == nil {
		return []map[string]any{}, errors.New("querier (db/tx/conn) is nil")
	}

	listMap := []map[string]any{}

	err := sqlscan.Select(ctx, q, &listMap, sqlStr, args...)

	if err != nil {
		if sqlscan.NotFound(err) {
			return []map[string]any{}, nil
		}

		return []map[string]any{}, err
	}

	return listMap, nil
}

func SelectToMapString(ctx context.Context, q Querier, sqlStr string, args ...any) ([]map[string]string, error) {
	if q == nil {
		return []map[string]string{}, errors.New("querier (db/tx/conn) is nil")
	}

	listMapAny, err := SelectToMapAny(ctx, q, sqlStr, args...)

	if err != nil {
		return []map[string]string{}, err
	}

	listMapString := []map[string]string{}

	for i := 0; i < len(listMapAny); i++ {
		mapString := maputils.MapStringAnyToMapStringString(listMapAny[i])
		listMapString = append(listMapString, mapString)
	}

	return listMapString, nil
}
