package database

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
)

func stmtMap(stmt *sql.Stmt, args ...any) (map[string]any, error) {
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))

	for i := range values {
		values[i] = new(any)
	}

	result := make(map[string]any, len(cols))

	if rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, err
		}

		for ii, key := range cols {
			if values[ii] == nil {
				continue
			}

			result[key] = *(values[ii].(*any))
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func stmtMapCtx(ctx context.Context, stmt *sql.Stmt, args ...any) (map[string]any, error) {
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))

	for i := range values {
		values[i] = new(any)
	}

	result := make(map[string]any, len(cols))

	if rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, err
		}

		for ii, key := range cols {
			if values[ii] == nil {
				continue
			}
			result[key] = *(values[ii].(*any))
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func stmtMapSlice(stmt *sql.Stmt, args ...any) ([]map[string]any, error) {
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))
	results := make([]map[string]any, 0)

	for i := range values {
		values[i] = new(any)
	}

	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, err
		}
		result := make(map[string]any, len(cols))
		for ii, key := range cols {
			if values[ii] == nil {
				continue
			}
			result[key] = *(values[ii].(*any))
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func stmtMapSliceCtx(ctx context.Context, stmt *sql.Stmt, args ...any) ([]map[string]any, error) {
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]any, len(cols))
	results := make([]map[string]any, 0)

	for i := range values {
		values[i] = new(any)
	}

	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, err
		}
		result := make(map[string]any, len(cols))
		for ii, key := range cols {
			if values[ii] == nil {
				continue
			}
			result[key] = *(values[ii].(*any))
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// 将mysql的占位符转换为postgres的占位符
func Rebind(query string) string {
	// 使用正则表达式匹配所有的问号
	re := regexp.MustCompile(`\?`)
	index := 1

	// 使用替换函数来替换每个问号
	result := re.ReplaceAllStringFunc(query, func(_ string) string {
		placeholder := fmt.Sprintf("$%d", index)
		index++
		return placeholder
	})

	return result
}
