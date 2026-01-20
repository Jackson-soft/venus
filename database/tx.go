package database

import (
	"context"
	"database/sql"
)

type Tx struct {
	tx_       *sql.Tx
	hasError_ bool // 有一些错误 - -
}

func (t *Tx) Close() error {
	if t.hasError_ {
		return t.tx_.Rollback()
	}

	err := t.tx_.Commit()
	if err != nil {
		return t.tx_.Rollback()
	}

	return nil
}

func (t *Tx) HasError() {
	t.hasError_ = true
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

func (t *Tx) ExecContext(ctx context.Context, query string, args ...any) (int64, error) {
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

func (t *Tx) QueryMapContext(ctx context.Context, query string, args ...any) (map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return stmtMapCtx(ctx, stmt, args...)
}

func (t *Tx) QueryMapSliceContext(ctx context.Context, query string, args ...any) ([]map[string]any, error) {
	stmt, err := t.tx_.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return stmtMapSliceCtx(ctx, stmt, args...)
}
