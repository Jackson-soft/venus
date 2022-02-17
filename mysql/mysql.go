package mysql

import (
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

func (m *MySQL) Close() error {
	err := m.conn_.Close()
	return err
}

func (m *MySQL) GetConn() *sql.DB {
	return m.conn_
}

func (m *MySQL) BeginTx() (*Tx, error) {
	tx, err := m.conn_.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return &Tx{tx_: tx, hasError_: false}, nil
}

func (m *MySQL) Insert(query string, args ...interface{}) (int64, error) {
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

func (m *MySQL) Delete(query string, args ...interface{}) (int64, error) {
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

func (m *MySQL) Update(query string, args ...interface{}) (int64, error) {
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

func (m *MySQL) QueryForMap(query string, args ...interface{}) (map[string]interface{}, error) {
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

	values := make([]interface{}, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	result := make(map[string]interface{}, len(cols))

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

func (m *MySQL) QueryForMapSlice(query string, args ...interface{}) ([]map[string]interface{}, error) {
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

	values := make([]interface{}, len(cols))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		if err = rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		result := make(map[string]interface{}, len(cols))
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
