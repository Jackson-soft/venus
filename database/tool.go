package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

// 将mysql的占位符转换为postgres的占位符
func Rebind(query string) string {
	var b strings.Builder

	n := 1

	for _, ch := range query {
		if ch == '?' {
			fmt.Fprintf(&b, "$%d", n)
			n++
		} else {
			b.WriteRune(ch)
		}
	}

	return b.String()
}

func rowMap(rows *sql.Rows) (map[string]any, error) {
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))
	ptrs := make([]any, len(cols))

	for i := range values {
		ptrs[i] = &values[i]
	}

	result := make(map[string]any, len(cols))

	if rows.Next() {
		err = rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}

		for ii, key := range cols {
			if b, ok := values[ii].([]byte); ok {
				result[key] = string(b)
			} else {
				result[key] = values[ii]
			}
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func rowMapSlice(rows *sql.Rows) ([]map[string]any, error) {
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))

	var results []map[string]any

	ptrs := make([]any, len(cols))
	for i := range values {
		ptrs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(ptrs...)
		if err != nil {
			return nil, err
		}

		result := make(map[string]any, len(cols))
		for ii, key := range cols {
			if b, ok := values[ii].([]byte); ok {
				result[key] = string(b)
			} else {
				result[key] = values[ii]
			}
		}

		results = append(results, result)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func stmtMapSliceCtx(ctx context.Context, stmt *sql.Stmt, args ...any) ([]map[string]any, error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rowMapSlice(rows)
}
