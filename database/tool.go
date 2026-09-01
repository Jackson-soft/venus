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

func StmtMapSliceCtx(ctx context.Context, stmt *sql.Stmt, args ...any) ([]map[string]any, error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rowMapSlice(rows)
}

func StmtMapCtx(ctx context.Context, stmt *sql.Stmt, args ...any) (map[string]any, error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rowMap(rows)
}

// rowScanner 封装一次查询的列名与扫描缓冲区，供 rowMap 与 rowMapSlice 复用。
type rowScanner struct {
	cols []string
	vals []any
	ptrs []any
}

// newRowScanner 读取列名并初始化扫描缓冲区。
func newRowScanner(rows *sql.Rows) (*rowScanner, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	vals := make([]any, len(cols))
	ptrs := make([]any, len(cols))

	for i := range vals {
		ptrs[i] = &vals[i]
	}

	return &rowScanner{cols: cols, vals: vals, ptrs: ptrs}, nil
}

// scan 扫描当前行，并将其转换为以列名为键的 map。
func (r *rowScanner) scan(rows *sql.Rows) (map[string]any, error) {
	err := rows.Scan(r.ptrs...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]any, len(r.cols))
	for i, key := range r.cols {
		result[key] = cellValue(r.vals[i])
	}

	return result, nil
}

// cellValue 将 SQL 扫描得到的 []byte 转为 string；其余类型（含 NULL 的 nil）原样返回。
func cellValue(v any) any {
	if b, ok := v.([]byte); ok {
		return string(b)
	}

	return v
}

// rowMap 读取单行结果并返回以列名为键的 map；无匹配行时返回 nil, nil。
func rowMap(rows *sql.Rows) (map[string]any, error) {
	defer rows.Close()

	scanner, err := newRowScanner(rows)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		err = rows.Err()
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	result, err := scanner.scan(rows)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// rowMapSlice 读取多行结果并返回以列名为键的 map 切片。
func rowMapSlice(rows *sql.Rows) ([]map[string]any, error) {
	defer rows.Close()

	scanner, err := newRowScanner(rows)
	if err != nil {
		return nil, err
	}

	var results []map[string]any

	for rows.Next() {
		var result map[string]any

		result, err = scanner.scan(rows)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil
}
