package mysql

import "time"

type Option func(*MySQL)

func WithMaxConns(count int) Option {
	return func(sql *MySQL) {
		sql.conn_.SetMaxOpenConns(count)
	}
}

func WithMaxIdle(count int) Option {
	return func(sql *MySQL) {
		sql.conn_.SetMaxIdleConns(count)
	}
}

func WithMaxLifetime(ex time.Duration) Option {
	return func(sql *MySQL) {
		sql.conn_.SetConnMaxLifetime(ex)
	}
}
