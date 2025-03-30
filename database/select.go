package database

import (
	"errors"

	"github.com/gouniverse/maputils"
)

func SelectToMapAny(ctx QueryableContext, sqlStr string, args ...any) ([]map[string]any, error) {
	if ctx.queryable == nil {
		return []map[string]any{}, errors.New("querier (db/tx/conn) is nil")
	}

	listMap := []map[string]any{}

	rows, err := ctx.queryable.QueryContext(ctx, sqlStr, args...)

	if err != nil {
		return []map[string]any{}, err
	}

	for rows.Next() {
		var row map[string]any
		err := rows.Scan(&row)
		if err != nil {
			return []map[string]any{}, err
		}
		listMap = append(listMap, row)
	}

	if err := rows.Err(); err != nil {
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
