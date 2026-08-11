package database

import (
	"context"
	"database/sql"
	"sync"
)

// 标准库的数据库简单封装

type Database struct {
	mu_   sync.RWMutex
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

	for _, op := range ops {
		op(database)
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
	return d.conn().PingContext(ctx)
}

func (d *Database) Close() error {
	return d.conn().Close()
}

func (d *Database) Client() *sql.DB {
	return d.conn()
}

func (d *Database) Reset(db *sql.DB, dsn string) {
	if db == nil {
		return
	}

	d.mu_.Lock()
	defer d.mu_.Unlock()

	d.conn_ = db
	d.dsn_ = dsn
}

func (d *Database) InsertContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := d.conn().ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...any) (int64, error) {
	res, err := d.conn().ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (d *Database) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	rows, err := d.conn().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rowMap(rows)
}

func (d *Database) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	rows, err := d.conn().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rowMapSlice(rows)
}

func (d *Database) conn() *sql.DB {
	d.mu_.RLock()
	defer d.mu_.RUnlock()

	return d.conn_
}
