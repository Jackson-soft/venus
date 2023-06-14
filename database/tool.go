package database

import (
	"database/sql"
	"reflect"
)

func stmtMap(stmt *sql.Stmt, args ...any) (map[string]any, error) {
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
