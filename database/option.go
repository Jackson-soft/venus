package database

import (
	"time"
)

type Option func(*Database)

func WithMaxConns(count int) Option {
	return func(db *Database) {
		db.conn_.SetMaxOpenConns(count)
	}
}

func WithMaxIdle(count int) Option {
	return func(db *Database) {
		db.conn_.SetMaxIdleConns(count)
	}
}

func WithMaxLifetime(ex time.Duration) Option {
	return func(db *Database) {
		db.conn_.SetConnMaxLifetime(ex)
	}
}

func WithMaxIdleTime(ex time.Duration) Option {
	return func(db *Database) {
		db.conn_.SetConnMaxIdleTime(ex)
	}
}
