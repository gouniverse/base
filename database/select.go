package database

import (
	"errors"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/gouniverse/maputils"
)

func SelectToMapAny(ctx QueryableContext, sqlStr string, args ...any) ([]map[string]any, error) {
	if ctx.queryable == nil {
		return []map[string]any{}, errors.New("querier (db/tx/conn) is nil")
	}

	listMap := []map[string]any{}

	err := sqlscan.Select(ctx, ctx.queryable, &listMap, sqlStr, args...)

	if err != nil {
		if sqlscan.NotFound(err) {
			return []map[string]any{}, nil
		}

		return []map[string]any{}, err
	}

	return listMap, nil
}

func SelectToMapString(ctx QueryableContext, sqlStr string, args ...any) ([]map[string]string, error) {
	if ctx.queryable == nil {
		return []map[string]string{}, errors.New("querier (db/tx/conn) is nil")
	}

	listMapAny, err := SelectToMapAny(ctx, sqlStr, args...)

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
