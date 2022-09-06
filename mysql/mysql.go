package mysql

import (
	"context"
	"database/sql"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	conn_ *sql.DB
	dsn_  string
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func (m *MySQL) Open(dsn string, ops ...Option) error {
	m.dsn_ = dsn
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	m.conn_ = conn

	if len(ops) > 0 {
		for _, op := range ops {
			op(m)
		}
	}
	return nil
}

func (m *MySQL) Ping(ctx context.Context) error {
	return m.conn_.PingContext(ctx)
}

func (m *MySQL) Close() error {
	err := m.conn_.Close()
	return err
}

func (m *MySQL) GetConn() *sql.DB {
	return m.conn_
}

func (m *MySQL) Reset(dsn string, db *sql.DB) {
	m.conn_ = db
	m.dsn_ = dsn
}

func (m *MySQL) BeginTx() (*Tx, error) {
	tx, err := m.conn_.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx_: tx, hasError_: false}, nil
}

func (m *MySQL) Insert(query string, args ...any) (int64, error) {
	stmt, err := m.conn_.Prepare(query)
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

func (m *MySQL) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := m.conn_.PrepareContext(ctx, query)
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

func (m *MySQL) Delete(query string, args ...any) (int64, error) {
	stmt, err := m.conn_.Prepare(query)
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

func (m *MySQL) Update(query string, args ...any) (int64, error) {
	stmt, err := m.conn_.Prepare(query)
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

func (m *MySQL) UpdateContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := m.conn_.PrepareContext(ctx, query)
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

func (m *MySQL) QueryForMap(query string, args ...any) (map[string]any, error) {
	stmt, err := m.conn_.Prepare(query)
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

func (m *MySQL) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	stmt, err := m.conn_.PrepareContext(ctx, query)
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

func (m *MySQL) QueryForMapSlice(query string, args ...any) ([]map[string]any, error) {
	stmt, err := m.conn_.Prepare(query)
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

	results := make([]map[string]any, 0)
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

func (m *MySQL) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	stmt, err := m.conn_.PrepareContext(ctx, query)
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

	results := make([]map[string]any, 0)
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
