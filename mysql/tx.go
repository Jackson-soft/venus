package mysql

import (
	"context"
	"database/sql"
	"reflect"
)

type Tx struct {
	tx_       *sql.Tx
	hasError_ bool // 有一些错误 - -
}

func (t *Tx) Close() error {
	if t.hasError_ {
		return t.tx_.Rollback()
	}

	if err := t.tx_.Commit(); err != nil {
		return t.tx_.Rollback()
	}
	return nil
}

func (t *Tx) HasError() {
	t.hasError_ = true
}

func (t *Tx) Insert(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (t *Tx) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func (t *Tx) Delete(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) DeleteContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) Update(query string, args ...any) (int64, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) UpdateContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (t *Tx) QueryForMap(query string, args ...any) (map[string]any, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return nil, err
	}

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

	scanArgs := make([]any, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := make(map[string]any, len(cols))

	if rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		for ii, key := range cols {
			if scanArgs[ii] == nil {
				continue
			}
			value := reflect.Indirect(reflect.ValueOf(scanArgs[ii]))
			if value.Elem().Kind() == reflect.Slice {
				result[key] = string(value.Interface().([]byte))
			} else {
				result[key] = value.Interface()
			}
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *Tx) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

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

	scanArgs := make([]any, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := make(map[string]any, len(cols))

	if rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}

		for ii, key := range cols {
			if scanArgs[ii] == nil {
				continue
			}
			value := reflect.Indirect(reflect.ValueOf(scanArgs[ii]))
			if value.Elem().Kind() == reflect.Slice {
				result[key] = string(value.Interface().([]byte))
			} else {
				result[key] = value.Interface()
			}
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (t *Tx) QueryForMapSlice(query string, args ...any) ([]map[string]any, error) {
	stmt, err := t.tx_.Prepare(query)
	if err != nil {
		return nil, err
	}
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

	scanArgs := make([]any, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var results []map[string]any
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		result := make(map[string]any, len(cols))
		for ii, key := range cols {
			if scanArgs[ii] == nil {
				continue
			}
			value := reflect.Indirect(reflect.ValueOf(scanArgs[ii]))
			if value.Elem().Kind() == reflect.Slice {
				result[key] = string(value.Interface().([]byte))
			} else {
				result[key] = value.Interface()
			}
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (t *Tx) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
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

	scanArgs := make([]any, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var results []map[string]any
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		result := make(map[string]any, len(cols))
		for ii, key := range cols {
			if scanArgs[ii] == nil {
				continue
			}
			value := reflect.Indirect(reflect.ValueOf(scanArgs[ii]))
			if value.Elem().Kind() == reflect.Slice {
				result[key] = string(value.Interface().([]byte))
			} else {
				result[key] = value.Interface()
			}
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
