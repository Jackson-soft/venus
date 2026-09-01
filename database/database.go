package database

import (
	"context"
	"database/sql"
	"errors"
	"sync"
)

// 标准库的数据库简单封装

type Database struct {
	mutex_  sync.RWMutex
	conn_   *sql.DB
	name_   string
	dsn_    string
	closed_ bool
}

func OpenDB(driverName, dsn string, ops ...Option) (*Database, error) {
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	database := &Database{
		mutex_:  sync.RWMutex{},
		conn_:   conn,
		name_:   driverName,
		dsn_:    dsn,
		closed_: false,
	}

	for _, op := range ops {
		op(database)
	}

	return database, nil
}

// NewDB 使用已存在的 *sql.DB 构造 Database，并可选地应用连接池 Option。
// 传入 nil 的 db 会 panic。
func NewDB(driverName string, db *sql.DB, ops ...Option) (*Database, error) {
	if db == nil {
		return nil, errors.New("database: NewDB: nil *sql.DB")
	}

	database := &Database{
		mutex_:  sync.RWMutex{},
		conn_:   db,
		name_:   driverName,
		closed_: false,
	}

	for _, op := range ops {
		op(database)
	}

	return database, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.conn().PingContext(ctx)
}

func (d *Database) Close() error {
	d.mutex_.Lock()
	if d.closed_ {
		d.mutex_.Unlock()

		return nil
	}

	d.closed_ = true
	conn := d.conn_
	d.mutex_.Unlock()

	return conn.Close()
}

func (d *Database) Client() *sql.DB {
	return d.conn()
}

// Reset 将内部连接替换为新连接，并复位关闭状态。
// 返回旧连接，由调用方决定何时关闭；不要在仍有 goroutine 使用时关闭旧连接。
func (d *Database) Reset(db *sql.DB, dsn string) *sql.DB {
	if db == nil {
		return nil
	}

	d.mutex_.Lock()
	defer d.mutex_.Unlock()

	old := d.conn_
	d.conn_ = db
	d.dsn_ = dsn
	d.closed_ = false

	return old
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

func (d *Database) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return d.conn().PrepareContext(ctx, query)
}

func (d *Database) conn() *sql.DB {
	d.mutex_.RLock()
	defer d.mutex_.RUnlock()

	return d.conn_
}
