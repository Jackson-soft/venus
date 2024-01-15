package database

import (
	"context"
	"database/sql"
)

// 标准库的数据库简单封装

type Database struct {
	conn_ *sql.DB
	name_ string
	dsn_  string
}

func OpenDB(driverName, dsn string, ops ...Option) (*Database, error) {
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	database := &Database{
		conn_: conn,
		name_: driverName,
		dsn_:  dsn,
	}

	if len(ops) > 0 {
		for _, op := range ops {
			op(database)
		}
	}
	return database, nil
}

func NewDB(driverName string, db *sql.DB) *Database {
	client := new(Database)
	client.conn_ = db
	client.name_ = driverName

	return client
}

func (d *Database) Ping(ctx context.Context) error {
	return d.conn_.PingContext(ctx)
}

func (d *Database) Close() error {
	return d.conn_.Close()
}

func (d *Database) Client() *sql.DB {
	return d.conn_
}

func (d *Database) Reset(db *sql.DB, dsn string) {
	if db == nil {
		return
	}
	d.conn_ = db
	d.dsn_ = dsn
}

func (d *Database) BeginTx() (*Tx, error) {
	tx, err := d.conn_.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{
		tx_:       tx,
		hasError_: false,
	}, nil
}

func (d *Database) BeginTxCtx(ctx context.Context) (*Tx, error) {
	tx, err := d.conn_.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Tx{
		tx_:       tx,
		hasError_: false,
	}, nil
}

func (d *Database) Insert(query string, args ...any) (int64, error) {
	stmt, err := d.conn_.Prepare(query)
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

func (d *Database) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := d.conn_.PrepareContext(ctx, query)
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

func (d *Database) Delete(query string, args ...any) (int64, error) {
	stmt, err := d.conn_.Prepare(query)
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

func (d *Database) Update(query string, args ...any) (int64, error) {
	stmt, err := d.conn_.Prepare(query)
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

func (d *Database) UpdateContext(ctx context.Context, query string, args ...any) (int64, error) {
	stmt, err := d.conn_.PrepareContext(ctx, query)
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

func (d *Database) QueryForMap(query string, args ...any) (map[string]any, error) {
	stmt, err := d.conn_.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMap(stmt, args...)
}

func (d *Database) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	stmt, err := d.conn_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMapCtx(ctx, stmt, args...)
}

func (d *Database) QueryForMapSlice(query string, args ...any) ([]map[string]any, error) {
	stmt, err := d.conn_.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMapSlice(stmt, args...)
}

func (d *Database) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	stmt, err := d.conn_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmtMapSliceCtx(ctx, stmt, args...)
}
