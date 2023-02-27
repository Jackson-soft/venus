package database

import "time"

type Option func(*Database)

func WithMaxConns(count int) Option {
	return func(sql *Database) {
		sql.conn_.SetMaxOpenConns(count)
	}
}

func WithMaxIdle(count int) Option {
	return func(sql *Database) {
		sql.conn_.SetMaxIdleConns(count)
	}
}

func WithMaxLifetime(ex time.Duration) Option {
	return func(sql *Database) {
		sql.conn_.SetConnMaxLifetime(ex)
	}
}
